package add

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
	"github.com/sirupsen/logrus"
)

func addItemToProject(ctx context.Context, logE *logrus.Entry, _ GitHub, item *Item, projectID string) error {
	var v4Client *githubv4.Client
	return addProjectNextItem(ctx, logE, v4Client, item.ID, projectID)
}

func addProjectNextItem(ctx context.Context, logE *logrus.Entry, v4Client *githubv4.Client, issueID, projectID string) error {
	var m struct {
		AddProjectNextItem struct {
			ProjectNextItem struct {
				Title string
			}
		} `graphql:"addProjectNextItem(input: $input)"`
	}
	logE.WithFields(logrus.Fields{
		"project_id": projectID,
		"content_id": issueID,
	}).Info("adding an issue or a pull request to a GitHub Project")
	if err := v4Client.Mutate(ctx, &m, githubv4.AddProjectV2ItemByIdInput{
		ProjectID: projectID,
		ContentID: issueID,
	}, nil); err != nil {
		return fmt.Errorf("add an issue or a pull request to a GitHub Project by GitHub GraphQL API (addProjectNextItem): %w", err)
	}
	return nil
}
