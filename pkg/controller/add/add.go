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

type Entry struct{}

type Item struct{}

type GitHub interface{}

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

func handleEntry(ctx context.Context, gh GitHub, _ *Config, _ *Entry) error {
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
		if err := addItemToProject(ctx, gh, item); err != nil {
			return fmt.Errorf("add an item to a project: %w", err)
		}
	}
	return nil
}

func findAndReadConfig(_ afero.Fs, _ *Config) error {
	return nil
}

func searchIssuesAndPRs(_ context.Context, _ GitHub) ([]*Item, error) {
	return nil, nil
}

// excludeItem returns true if the item should be excluded.
func excludeItem(_ *Item) bool {
	return false
}

func addItemToProject(_ context.Context, _ GitHub, _ *Item) error {
	return nil
}
