package cli

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/ghproj/pkg/controller/add"
	"github.com/suzuki-shunsuke/ghproj/pkg/log"
	"github.com/urfave/cli/v2"
)

type addCommand struct {
	logE *logrus.Entry
}

func (rc *addCommand) command() *cli.Command {
	return &cli.Command{
		Name:      "add",
		Usage:     "Add GitHub Issues and Pull Requests to GitHub Projects",
		UsageText: "ghproj set",
		Description: `Add GitHub Issues and Pull Requests to GitHub Projects.

$ ghproj add
`,
		Action: rc.action,
		Flags:  []cli.Flag{},
	}
}

func (rc *addCommand) action(c *cli.Context) error {
	fs := afero.NewOsFs()
	logE := rc.logE
	log.SetLevel(c.String("log-level"), logE)
	log.SetColor(c.String("log-color"), logE)
	return add.Add(c.Context, logE, fs, &add.Param{ //nolint:wrapcheck
	})
}
