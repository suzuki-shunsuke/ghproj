package github

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
)

type Item struct {
	ID     string // issue and pull request id
	Title  string
	Labels []string
	Org    string
	User   string
	Repo   string
	Open   bool
	Merged bool
}

func (c *Client) SearchItems(ctx context.Context, query string) ([]*Item, error) {
	var q struct {
		Search struct {
			Nodes []struct {
				Issue struct {
					ID    githubv4.String
					Title githubv4.String
				} `graphql:"... on Issue"`
				PullRequest struct {
					ID    githubv4.String
					Title githubv4.String
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
			item := &Item{
				ID:    string(node.Issue.ID),
				Title: string(node.Issue.Title),
				// Number: int(node.Issue.Number),
			}
			if node.PullRequest.ID != "" {
				item = &Item{
					ID:    string(node.PullRequest.ID),
					Title: string(node.PullRequest.Title),
				}
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
