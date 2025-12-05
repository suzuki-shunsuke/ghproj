package cli

import (
	"context"
	"log/slog"
	"os"

	"github.com/suzuki-shunsuke/go-stdutil"
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/urfave"
	"github.com/urfave/cli/v3"
)

func Run(ctx context.Context, logger *slog.Logger, logLevelVar *slog.LevelVar, ldFlags *stdutil.LDFlags, args ...string) error {
	return urfave.Command(ldFlags, &cli.Command{ //nolint:wrapcheck
		Name:    "ghproj",
		Usage:   "",
		Version: ldFlags.Version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "log-level",
				Usage: "log level",
			},
		},
		EnableShellCompletion: true,
		Commands: []*cli.Command{
			(&initCommand{
				logger:      logger,
				logLevelVar: logLevelVar,
			}).command(),
			(&addCommand{
				logger:      logger,
				logLevelVar: logLevelVar,
			}).command(),
			(&completionCommand{
				stdout: os.Stdout,
			}).command(),
		},
	}).Run(ctx, args)
}
