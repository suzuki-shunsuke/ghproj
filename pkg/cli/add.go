package cli

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/ghproj/pkg/controller/add"
	"github.com/suzuki-shunsuke/ghproj/pkg/github"
	"github.com/suzuki-shunsuke/ghproj/pkg/log"
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
				EnvVars: []string{"GHPROJ_CONFIG"},
			},
		},
	}
}

func (rc *addCommand) action(c *cli.Context) error {
	fs := afero.NewOsFs()
	logE := rc.logE
	log.SetLevel(c.String("log-level"), logE)
	log.SetColor(c.String("log-color"), logE)
	gh := github.New(c.Context, os.Getenv("GITHUB_TOKEN"))
	return add.Add(c.Context, logE, fs, gh, &add.Param{ //nolint:wrapcheck
		ConfigFilePath: c.String("config"),
		ConfigText:     os.Getenv("GHPROJ_CONFIG_TEXT"),
	})
}
