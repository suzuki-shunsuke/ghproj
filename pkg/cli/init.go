package cli

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/ghproj/pkg/controller/initcmd"
	"github.com/suzuki-shunsuke/ghproj/pkg/log"
	"github.com/urfave/cli/v3"
)

type initCommand struct {
	logE *logrus.Entry
}

func (rc *initCommand) command() *cli.Command {
	return &cli.Command{
		Name:      "init",
		Usage:     "Scaffold a configuration file",
		UsageText: "ghproj init",
		Description: `Scaffold a configuration file.

$ ghproj init
`,
		Action: rc.action,
		Flags:  []cli.Flag{},
	}
}

func (rc *initCommand) action(c *cli.Context) error {
	fs := afero.NewOsFs()
	logE := rc.logE
	log.SetLevel(c.String("log-level"), logE)
	log.SetColor(c.String("log-color"), logE)
	ctrl := initcmd.NewController(fs)
	return ctrl.Init(c.Context, logE) //nolint:wrapcheck
}
