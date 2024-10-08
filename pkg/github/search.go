package github

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
)

type Item struct {
	ID         string // issue and pull request id
	State      string
	Title      string
	Labels     []string
	Open       bool
	Merged     bool
	IsArchived bool
	Repo       Repo
	Author     string
}

type Repo struct {
	Owner      string
	Name       string
	IsArchived bool
	IsFork     bool
}

func (c *Client) SearchItems(ctx context.Context, query string) ([]*Item, error) { //nolint:funlen
	var q struct {
		Search struct {
			Nodes []struct {
				Issue struct {
					ID         githubv4.String
					Title      githubv4.String
					Repository Repository
				} `graphql:"... on Issue"`
				PullRequest struct {
					ID         githubv4.String
					Title      githubv4.String
					Repository Repository
				} `graphql:"... on PullRequest"`
			}
			PageInfo struct {
				EndCursor   githubv4.String
				HasNextPage bool
			}
		} `graphql:"search(first: 100, query: $searchQuery, type: $searchType, after: $cursor)"`
	}
	variables := map[string]interface{}{
		"searchQuery": githubv4.String(query),
		"searchType":  githubv4.SearchTypeIssue,
		"cursor":      (*githubv4.String)(nil),
	}
	var items []*Item
	for range 30 {
		if err := c.v4Client.Query(ctx, &q, variables); err != nil {
			return nil, fmt.Errorf("get an issue by GitHub GraphQL API: %w", err)
		}
		for _, node := range q.Search.Nodes {
			var item *Item
			switch {
			case node.Issue.Title != "":
				item = &Item{
					ID:    string(node.Issue.ID),
					Title: string(node.Issue.Title),
					Repo: Repo{
						Owner:      node.Issue.Repository.Owner.Login,
						Name:       node.Issue.Repository.Name,
						IsArchived: node.Issue.Repository.IsArchived,
						IsFork:     node.Issue.Repository.IsFork,
					},
				}
			case node.PullRequest.Title != "":
				item = &Item{
					ID:    string(node.PullRequest.ID),
					Title: string(node.PullRequest.Title),
					Repo: Repo{
						Owner:      node.PullRequest.Repository.Owner.Login,
						Name:       node.PullRequest.Repository.Name,
						IsArchived: node.PullRequest.Repository.IsArchived,
						IsFork:     node.PullRequest.Repository.IsFork,
					},
				}
			default:
				continue
			}
			items = append(items, item)
		}

		if !q.Search.PageInfo.HasNextPage {
			return items, nil
		}
		variables["cursor"] = githubv4.NewString(q.Search.PageInfo.EndCursor)
	}
	return items, nil
}
