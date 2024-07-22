package github

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
)

/*
query {
  node (id: "PVT_kwHOAMtMJ84AQCf4") {
    ... on ProjectV2 {
      number
      title
      items(first: 100) {
        pageInfo {
          endCursor
          hasNextPage
          startCursor
        }
        nodes {
          type
          id
          content {
            ... on Issue {
              state
              closed
              id
              author {
                login
              }
              labels(first: 100) {
                nodes {
                  name
                }
              }
              repository {
                nameWithOwner
                name
                owner {
                  login
                }
              }
            }
            ... on PullRequest {
              prState: state
              merged
              closed
              id
              author {
                login
              }
              repository {
                nameWithOwner
                name
                owner {
                  login
                }
              }
            }
          }
          fieldValueByName( name: "Status") {
             ... on ProjectV2ItemFieldSingleSelectValue {
              status: name
            }
          }
        }
      }
    }
  }
}
*/

type Repository struct {
	NameWithOwner string
	Name          string
	IsArchived    bool
	IsFork        bool
	Owner         struct {
		Login string
	}
}

type Labels struct {
	Nodes []struct {
		Name string
	}
}

type Author struct {
	Login string
}

func (c *Client) ListItems(ctx context.Context, projectID string) ([]*Item, error) { //nolint:funlen,cyclop
	var q struct {
		Node struct {
			ProjectV2 struct {
				Items struct {
					Nodes []struct {
						ID         string
						Type       string
						IsArchived bool
						Content    struct {
							Issue struct {
								State      string
								Closed     bool
								ID         string
								Title      string
								Author     Author
								Labels     Labels `graphql:"labels(first: 100)"`
								Repository Repository
							} `graphql:"... on Issue"`
							PullRequest struct {
								State      string
								Merged     bool
								Closed     bool
								ID         string
								Title      string
								Author     Author
								Labels     Labels `graphql:"labels(first: 100)"`
								Repository Repository
							} `graphql:"... on PullRequest"`
						}
					}
					PageInfo struct {
						EndCursor   githubv4.String
						HasNextPage bool
					}
				} `graphql:"items(first: 100, after: $cursor)"`
			} `graphql:"... on ProjectV2"`
		} `graphql:"node(id: $id)"`
	}
	variables := map[string]interface{}{
		"id":     githubv4.ID(projectID),
		"cursor": (*githubv4.String)(nil),
	}
	var items []*Item
	for range 30 {
		if err := c.v4Client.Query(ctx, &q, variables); err != nil {
			return nil, fmt.Errorf("get GitHub Project items by GitHub GraphQL API: %w", err)
		}
		for _, node := range q.Node.ProjectV2.Items.Nodes {
			if node.IsArchived {
				continue
			}
			var item *Item
			switch node.Type {
			case "ISSUE":
				item = &Item{
					ID:     node.ID,
					State:  node.Content.Issue.State,
					Title:  node.Content.Issue.Title,
					Labels: make([]string, len(node.Content.Issue.Labels.Nodes)),
					Open:   !node.Content.Issue.Closed,
					Author: node.Content.Issue.Author.Login,
					Repo: Repo{
						Owner:      node.Content.Issue.Repository.Owner.Login,
						Repo:       node.Content.Issue.Repository.Name,
						IsArchived: node.Content.Issue.Repository.IsArchived,
						IsFork:     node.Content.Issue.Repository.IsFork,
					},
				}
				for i, label := range node.Content.Issue.Labels.Nodes {
					item.Labels[i] = label.Name
				}
			case "PULL_REQUEST":
				item = &Item{
					ID:     node.ID,
					State:  node.Content.PullRequest.State,
					Title:  node.Content.PullRequest.Title,
					Labels: make([]string, len(node.Content.Issue.Labels.Nodes)),
					Open:   !node.Content.PullRequest.Closed,
					Merged: !node.Content.PullRequest.Merged,
					Author: node.Content.PullRequest.Author.Login,
					Repo: Repo{
						Owner:      node.Content.Issue.Repository.Owner.Login,
						Repo:       node.Content.Issue.Repository.Name,
						IsArchived: node.Content.Issue.Repository.IsArchived,
						IsFork:     node.Content.Issue.Repository.IsFork,
					},
				}
				for i, label := range node.Content.PullRequest.Labels.Nodes {
					item.Labels[i] = label.Name
				}
			default:
				continue
			}
			items = append(items, item)
		}

		if !q.Node.ProjectV2.Items.PageInfo.HasNextPage {
			return items, nil
		}
		variables["cursor"] = githubv4.NewString(q.Node.ProjectV2.Items.PageInfo.EndCursor)
	}
	return items, nil
}
