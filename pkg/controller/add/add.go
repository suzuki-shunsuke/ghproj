package add

import (
	"context"
	"fmt"

	"github.com/spf13/afero"
)

type Param struct{}

type Config struct {
	Entries []*Entry
}

type Entry struct {
	ProjectID string
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

func Add(ctx context.Context, fs afero.Fs, _ *Param) error {
	cfg := &Config{}
	if err := findAndReadConfig(fs, cfg); err != nil {
		return fmt.Errorf("find and read a configuration file: %w", err)
	}
	var gh GitHub

	for _, entry := range cfg.Entries {
		if err := handleEntry(ctx, gh, cfg, entry); err != nil {
			return fmt.Errorf("handle an entry: %w", err)
		}
	}
	return nil
}

func handleEntry(ctx context.Context, gh GitHub, _ *Config, entry *Entry) error {
	// Search GitHub Issues and Pull Requests
	items, err := searchIssuesAndPRs(ctx, gh)
	if err != nil {
		return fmt.Errorf("search GitHub Issues and Pull Requests: %w", err)
	}
	// Add issues and pull requests to GitHub Projects
	for _, item := range items {
		// Exclude issues and pull requests based on the configuration
		if excludeItem(item) {
			continue
		}
		if err := addItemToProject(ctx, gh, item, entry.ProjectID); err != nil {
			return fmt.Errorf("add an item to a project: %w", err)
		}
	}
	return nil
}

// excludeItem returns true if the item should be excluded.
func excludeItem(_ *Item) bool {
	return false
}
