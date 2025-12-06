package cli

import (
	"context"
	"fmt"

	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/ghproj/pkg/controller/initcmd"
	"github.com/suzuki-shunsuke/slog-util/slogutil"
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/urfave"
	"github.com/urfave/cli/v3"
)

type initCommand struct{}

func (rc *initCommand) command(logger *slogutil.Logger) *cli.Command {
	return &cli.Command{
		Name:      "init",
		Usage:     "Scaffold a configuration file",
		UsageText: "ghproj init",
		Description: `Scaffold a configuration file.

$ ghproj init
`,
		Action: urfave.Action(rc.action, logger),
		Flags:  []cli.Flag{},
	}
}

func (rc *initCommand) action(ctx context.Context, c *cli.Command, logger *slogutil.Logger) error {
	fs := afero.NewOsFs()
	if err := logger.SetLevel(c.String("log-level")); err != nil {
		return fmt.Errorf("set log level: %w", err)
	}
	ctrl := initcmd.NewController(fs)
	return ctrl.Init(ctx, logger.Logger) //nolint:wrapcheck
}
