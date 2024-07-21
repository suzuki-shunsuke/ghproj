package add

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
)

type GitHub interface{}

func searchIssuesAndPRs(ctx context.Context, _ GitHub) ([]*Item, error) {
	var v4Client *githubv4.Client
	return listIssues(ctx, v4Client, "", "", "")
}

func listIssues(ctx context.Context, v4Client *githubv4.Client, repoOwner, repoName, title string) ([]*Item, error) {
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
		"searchQuery": githubv4.String(fmt.Sprintf(`repo:%s/%s state:open "%s" in:title`, repoOwner, repoName, title)),
		"searchType":  githubv4.SearchTypeIssue,
		"cursor":      (*githubv4.String)(nil),
	}
	var items []*Item
	for range 30 {
		if err := v4Client.Query(ctx, &q, variables); err != nil {
			return nil, fmt.Errorf("get an issue by GitHub GraphQL API: %w", err)
		}
		for _, node := range q.Search.Nodes {
			if title != string(node.Issue.Title) {
				continue
			}
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
