package add

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/vm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/ghproj/pkg/github"
)

type Param struct {
	ConfigFilePath string
	ConfigText     string
}

type Config struct {
	Entries []*Entry
}

type Entry struct {
	ProjectID string `yaml:"project_id"`
	Action    string
	Query     string
	Expr      string
	exprProg  *vm.Program
}

type GitHub interface {
	AddItemToProject(ctx context.Context, logE *logrus.Entry, input *github.InputAddItemToProject) error
	ArchiveItem(ctx context.Context, logE *logrus.Entry, input *github.InputArchiveItem) error
	SearchItems(ctx context.Context, query string) ([]*github.Item, error)
	ListItems(ctx context.Context, projectID string) ([]*github.Item, error)
}

func Add(ctx context.Context, logE *logrus.Entry, fs afero.Fs, gh GitHub, param *Param) error {
	cfg := &Config{}
	if err := findAndReadConfig(fs, cfg, param); err != nil {
		return fmt.Errorf("find and read a configuration file: %w", err)
	}
	if err := cfg.Validate(); err != nil {
		return fmt.Errorf("validate a configuration file: %w", err)
	}
	for _, entry := range cfg.Entries {
		if err := handleEntry(ctx, logE, gh, cfg, entry); err != nil {
			return fmt.Errorf("handle an entry: %w", err)
		}
	}
	return nil
}

func listItems(ctx context.Context, gh GitHub, entry *Entry) ([]*github.Item, error) {
	if entry.Action == "archive" {
		return gh.ListItems(ctx, entry.ProjectID) //nolint:wrapcheck
	}
	return gh.SearchItems(ctx, strings.ReplaceAll(entry.Query, "\n", " ")) //nolint:wrapcheck
}

func handleEntry(ctx context.Context, logE *logrus.Entry, gh GitHub, _ *Config, entry *Entry) error {
	// Search GitHub Issues and Pull Requests
	items, err := listItems(ctx, gh, entry)
	if err != nil {
		return fmt.Errorf("search GitHub Issues and Pull Requests: %w", err)
	}
	logE.WithFields(logrus.Fields{
		"number_of_items": len(items),
	}).Info("search issues and pull requests")
	// Add issues and pull requests to GitHub Projects
	for _, item := range items {
		// Exclude issues and pull requests based on the configuration
		if f, err := includeItem(item, entry); err != nil {
			return err
		} else if !f {
			continue
		}
		switch entry.Action {
		case "archive":
			if err := gh.ArchiveItem(ctx, logE, &github.InputArchiveItem{
				ProjectID: entry.ProjectID,
				ItemID:    item.ID,
			}); err != nil {
				return fmt.Errorf("archive an item: %w", err)
			}
		case "", "add":
			if err := gh.AddItemToProject(ctx, logE, &github.InputAddItemToProject{
				ProjectID: entry.ProjectID,
				ContentID: item.ID,
			}); err != nil {
				return fmt.Errorf("add an item to a project: %w", err)
			}
		}
	}
	return nil
}

// includeItem returns true if the item should be excluded.
func includeItem(item *github.Item, entry *Entry) (bool, error) {
	if entry.exprProg == nil {
		return true, nil
	}
	value, err := expr.Run(entry.exprProg, map[string]any{
		"Item": item,
	})
	if err != nil {
		return false, fmt.Errorf("evaluate the expression: %w", err)
	}
	b, ok := value.(bool)
	if !ok {
		return false, errors.New("the expression must return a boolean value")
	}
	return b, nil
}
