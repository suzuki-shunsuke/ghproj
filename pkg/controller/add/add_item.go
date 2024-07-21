package add

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
)

func addItemToProject(ctx context.Context, _ GitHub, item *Item, projectID string) error {
	var v4Client *githubv4.Client
	return addProjectNextItem(ctx, v4Client, item.ID, projectID)
}

func addProjectNextItem(ctx context.Context, v4Client *githubv4.Client, issueID, projectID string) error {
	var m struct {
		AddProjectNextItem struct {
			ProjectNextItem struct {
				Title string
			}
		} `graphql:"addProjectNextItem(input: $input)"`
	}
	if err := v4Client.Mutate(ctx, &m, githubv4.AddProjectV2ItemByIdInput{
		ProjectID: projectID,
		ContentID: issueID,
	}, nil); err != nil {
		return fmt.Errorf("add an issue to GitHub Project by GitHub GraphQL API (addProjectNextItem): %w", err)
	}
	return nil
}
