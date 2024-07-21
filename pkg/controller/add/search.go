package add

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
)

type GitHub interface{}

func searchIssuesAndPRs(ctx context.Context, _ GitHub, query string) ([]*Item, error) {
	var v4Client *githubv4.Client
	return listIssues(ctx, v4Client, query)
}

func listIssues(ctx context.Context, v4Client *githubv4.Client, query string) ([]*Item, error) {
	var q struct {
		Search struct {
			Nodes []struct {
				Issue struct {
					ID    githubv4.String
					Title githubv4.String
				} `graphql:"... on Issue"`
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
		if err := v4Client.Query(ctx, &q, variables); err != nil {
			return nil, fmt.Errorf("get an issue by GitHub GraphQL API: %w", err)
		}
		for _, node := range q.Search.Nodes {
			issue := &Item{
				ID:    string(node.Issue.ID),
				Title: string(node.Issue.Title),
				// Number: int(node.Issue.Number),
			}
			items = append(items, issue)
		}

		if !q.Search.PageInfo.HasNextPage {
			return items, nil
		}
		variables["cursor"] = githubv4.NewString(q.Search.PageInfo.EndCursor)
	}
	return items, nil
}
