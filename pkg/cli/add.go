package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/ghproj/pkg/controller/add"
	"github.com/suzuki-shunsuke/ghproj/pkg/github"
	"github.com/suzuki-shunsuke/slog-util/slogutil"
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/urfave"
	"github.com/urfave/cli/v3"
)

type addCommand struct{}

func (rc *addCommand) command(logger *slogutil.Logger) *cli.Command {
	return &cli.Command{
		Name:  "add",
		Usage: "Add GitHub Issues and Pull Requests to GitHub Projects",
		Description: `Add GitHub Issues and Pull Requests to GitHub Projects.

$ ghproj add
`,
		Action: urfave.Action(rc.action, logger),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "configuration file path",
				Sources: cli.EnvVars("GHPROJ_CONFIG"),
			},
		},
	}
}

func (rc *addCommand) action(ctx context.Context, c *cli.Command, logger *slogutil.Logger) error {
	fs := afero.NewOsFs()
	if err := logger.SetLevel(c.String("log-level")); err != nil {
		return fmt.Errorf("set log level: %w", err)
	}
	gh := github.New(ctx, os.Getenv("GITHUB_TOKEN"))
	return add.Add(ctx, logger.Logger, fs, gh, &add.Param{ //nolint:wrapcheck
		ConfigFilePath: c.String("config"),
		ConfigText:     os.Getenv("GHPROJ_CONFIG_TEXT"),
	})
}
