package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/ghproj/pkg/controller/add"
	"github.com/suzuki-shunsuke/ghproj/pkg/github"
	"github.com/suzuki-shunsuke/slog-util/slogutil"
	"github.com/urfave/cli/v3"
)

type addCommand struct{}

type addFlags struct {
	*GlobalFlags

	Config string
}

func (rc *addCommand) command(logger *slogutil.Logger, globalFlags *GlobalFlags) *cli.Command {
	flags := &addFlags{
		GlobalFlags: globalFlags,
	}
	return &cli.Command{
		Name:  "add",
		Usage: "Add GitHub Issues and Pull Requests to GitHub Projects",
		Description: `Add GitHub Issues and Pull Requests to GitHub Projects.

$ ghproj add
`,
		Action: func(ctx context.Context, _ *cli.Command) error {
			return rc.action(ctx, logger, flags)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "configuration file path",
				Sources:     cli.EnvVars("GHPROJ_CONFIG"),
				Destination: &flags.Config,
			},
		},
	}
}

func (rc *addCommand) action(ctx context.Context, logger *slogutil.Logger, flags *addFlags) error {
	fs := afero.NewOsFs()
	if err := logger.SetLevel(flags.LogLevel); err != nil {
		return fmt.Errorf("set log level: %w", err)
	}
	gh := github.New(ctx, os.Getenv("GITHUB_TOKEN"))
	return add.Add(ctx, logger.Logger, fs, gh, &add.Param{ //nolint:wrapcheck
		ConfigFilePath: flags.Config,
		ConfigText:     os.Getenv("GHPROJ_CONFIG_TEXT"),
	})
}
