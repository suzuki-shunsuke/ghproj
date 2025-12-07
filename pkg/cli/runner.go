package cli

import (
	"context"

	"github.com/suzuki-shunsuke/slog-util/slogutil"
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/urfave"
	"github.com/urfave/cli/v3"
)

type GlobalFlags struct {
	LogLevel string
}

func Run(ctx context.Context, logger *slogutil.Logger, env *urfave.Env) error {
	globalFlags := &GlobalFlags{}
	return urfave.Command(env, &cli.Command{ //nolint:wrapcheck
		Name:  "ghproj",
		Usage: "Add GitHub Issues and Pull Requests to GitHub Projects",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "log-level",
				Usage:       "log level",
				Destination: &globalFlags.LogLevel,
			},
		},
		Commands: []*cli.Command{
			(&initCommand{}).command(logger, globalFlags),
			(&addCommand{}).command(logger, globalFlags),
		},
	}).Run(ctx, env.Args)
}
