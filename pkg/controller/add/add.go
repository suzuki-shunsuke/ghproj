package add

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
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

func Add(ctx context.Context, logE *logrus.Entry, fs afero.Fs, _ *Param) error {
	cfg := &Config{}
	if err := findAndReadConfig(fs, cfg); err != nil {
		return fmt.Errorf("find and read a configuration file: %w", err)
	}
	var gh GitHub

	for _, entry := range cfg.Entries {
		if err := handleEntry(ctx, logE, gh, cfg, entry); err != nil {
			return fmt.Errorf("handle an entry: %w", err)
		}
	}
	return nil
}

func handleEntry(ctx context.Context, logE *logrus.Entry, gh GitHub, _ *Config, entry *Entry) error {
	// Search GitHub Issues and Pull Requests
	items, err := searchIssuesAndPRs(ctx, gh, entry.Query)
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
		if err := addItemToProject(ctx, logE, gh, item, entry.ProjectID); err != nil {
			return fmt.Errorf("add an item to a project: %w", err)
		}
	}
	return nil
}

// excludeItem returns true if the item should be excluded.
func excludeItem(_ *Item) bool {
	return false
}
