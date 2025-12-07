package cli

import (
	"context"
	"fmt"

	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/ghproj/pkg/controller/initcmd"
	"github.com/suzuki-shunsuke/slog-util/slogutil"
	"github.com/urfave/cli/v3"
)

type initCommand struct{}

func (rc *initCommand) command(logger *slogutil.Logger, globalFlags *GlobalFlags) *cli.Command {
	return &cli.Command{
		Name:      "init",
		Usage:     "Scaffold a configuration file",
		UsageText: "ghproj init",
		Description: `Scaffold a configuration file.

$ ghproj init
`,
		Action: func(ctx context.Context, _ *cli.Command) error {
			return rc.action(ctx, logger, globalFlags)
		},
	}
}

func (rc *initCommand) action(ctx context.Context, logger *slogutil.Logger, flags *GlobalFlags) error {
	fs := afero.NewOsFs()
	if err := logger.SetLevel(flags.LogLevel); err != nil {
		return fmt.Errorf("set log level: %w", err)
	}
	ctrl := initcmd.NewController(fs)
	return ctrl.Init(ctx, logger.Logger) //nolint:wrapcheck
}
