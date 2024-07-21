package add

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/ghproj/pkg/github"
)

type Param struct{}

type Config struct {
	Entries []*Entry
}

type Entry struct {
	ProjectID string `yaml:"project_id"`
	Query     string
}

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

type GitHub interface {
	AddItemToProject(ctx context.Context, logE *logrus.Entry, input *github.InputAddItemToProject) error
	SearchItems(ctx context.Context, query string) ([]*github.Item, error)
}

func Add(ctx context.Context, logE *logrus.Entry, fs afero.Fs, gh GitHub, _ *Param) error {
	cfg := &Config{}
	if err := findAndReadConfig(fs, cfg); err != nil {
		return fmt.Errorf("find and read a configuration file: %w", err)
	}
	for _, entry := range cfg.Entries {
		if err := handleEntry(ctx, logE, gh, cfg, entry); err != nil {
			return fmt.Errorf("handle an entry: %w", err)
		}
	}
	return nil
}

func handleEntry(ctx context.Context, logE *logrus.Entry, gh GitHub, _ *Config, entry *Entry) error {
	// Search GitHub Issues and Pull Requests
	items, err := gh.SearchItems(ctx, strings.ReplaceAll(entry.Query, "\n", " "))
	if err != nil {
		return fmt.Errorf("search GitHub Issues and Pull Requests: %w", err)
	}
	logE.WithFields(logrus.Fields{
		"number_of_items": len(items),
	}).Info("search issues and pull requests")
	// Add issues and pull requests to GitHub Projects
	for _, item := range items {
		// Exclude issues and pull requests based on the configuration
		if excludeItem(item) {
			continue
		}
		if err := gh.AddItemToProject(ctx, logE, &github.InputAddItemToProject{
			ProjectID: entry.ProjectID,
			ContentID: item.ID,
		}); err != nil {
			return fmt.Errorf("add an item to a project: %w", err)
		}
	}
	return nil
}

// excludeItem returns true if the item should be excluded.
func excludeItem(_ *github.Item) bool {
	return false
}
