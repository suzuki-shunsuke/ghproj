package github

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
	"github.com/sirupsen/logrus"
)

type InputArchiveItem struct {
	ItemID    string
	ProjectID string
}

func (c *Client) ArchiveItem(ctx context.Context, logE *logrus.Entry, input *InputArchiveItem) error {
	var m struct {
		ArchiveProjectItem struct {
			ProjectV2Item ProjectItem `graphql:"item"`
		} `graphql:"archiveProjectV2Item(input:$input)"`
	}
	logE.WithFields(logrus.Fields{
		"project_id": input.ProjectID,
		"item_id":    input.ItemID,
	}).Info("archiving a GitHub Project item")
	if err := c.v4Client.Mutate(ctx, &m, githubv4.ArchiveProjectV2ItemInput{
		ProjectID: input.ProjectID,
		ItemID:    input.ItemID,
	}, nil); err != nil {
		return fmt.Errorf("add a GitHub Project item by GitHub GraphQL API (archiveProjectV2Item): %w", err)
	}
	return nil
}
