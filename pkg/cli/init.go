package cli

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/ghproj/pkg/controller/initcmd"
	"github.com/suzuki-shunsuke/slog-util/slogutil"
	"github.com/urfave/cli/v3"
)

type initCommand struct {
	logger      *slog.Logger
	logLevelVar *slog.LevelVar
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
	logger := rc.logger
	if err := slogutil.SetLevel(rc.logLevelVar, c.String("log-level")); err != nil {
		return fmt.Errorf("set log level: %w", err)
	}
	ctrl := initcmd.NewController(fs)
	return ctrl.Init(ctx, logger) //nolint:wrapcheck
}
