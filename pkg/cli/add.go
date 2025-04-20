package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/ghproj/pkg/controller/add"
	"github.com/suzuki-shunsuke/ghproj/pkg/github"
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/log"
	"github.com/urfave/cli/v3"
)

type addCommand struct {
	logE *logrus.Entry
}

func (rc *addCommand) command() *cli.Command {
	return &cli.Command{
		Name:  "add",
		Usage: "Add GitHub Issues and Pull Requests to GitHub Projects",
		Description: `Add GitHub Issues and Pull Requests to GitHub Projects.

$ ghproj add
`,
		Action: rc.action,
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

func (rc *addCommand) action(ctx context.Context, c *cli.Command) error {
	fs := afero.NewOsFs()
	logE := rc.logE
	if err := log.Set(logE, c.String("log-level"), c.String("log-color")); err != nil {
		return fmt.Errorf("configure logger: %w", err)
	}
	gh := github.New(ctx, os.Getenv("GITHUB_TOKEN"))
	return add.Add(ctx, logE, fs, gh, &add.Param{ //nolint:wrapcheck
		ConfigFilePath: c.String("config"),
		ConfigText:     os.Getenv("GHPROJ_CONFIG_TEXT"),
	})
}
