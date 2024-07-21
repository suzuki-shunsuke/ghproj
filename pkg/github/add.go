package github

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
	"github.com/sirupsen/logrus"
)

type InputAddItemToProject struct {
	ContentID string
	ProjectID string
}

func (c *Client) AddItemToProject(ctx context.Context, logE *logrus.Entry, input *InputAddItemToProject) error {
	var m struct {
		CreateProjectItem struct {
			ProjectV2Item ProjectItem `graphql:"item"`
		} `graphql:"addProjectV2ItemById(input:$input)"`
	}
	logE.WithFields(logrus.Fields{
		"project_id": input.ProjectID,
		"content_id": input.ContentID,
	}).Info("adding an issue or a pull request to a GitHub Project")
	if err := c.v4Client.Mutate(ctx, &m, githubv4.AddProjectV2ItemByIdInput{
		ProjectID: input.ProjectID,
		ContentID: input.ContentID,
	}, nil); err != nil {
		return fmt.Errorf("add an issue or a pull request to a GitHub Project by GitHub GraphQL API (addProjectNextItem): %w", err)
	}
	return nil
}

type ProjectItem struct {
	Id string
}
