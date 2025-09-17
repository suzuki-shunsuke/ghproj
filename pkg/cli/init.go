package cli

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/ghproj/pkg/controller/initcmd"
	"github.com/suzuki-shunsuke/logrus-util/log"
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

func (rc *initCommand) action(ctx context.Context, c *cli.Command) error {
	fs := afero.NewOsFs()
	logE := rc.logE
	if err := log.Set(logE, c.String("log-level"), c.String("log-color")); err != nil {
		return fmt.Errorf("configure logger: %w", err)
	}
	ctrl := initcmd.NewController(fs)
	return ctrl.Init(ctx, logE) //nolint:wrapcheck
}
