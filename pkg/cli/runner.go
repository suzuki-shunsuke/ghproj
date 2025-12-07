package cli

import (
	"context"

	"github.com/suzuki-shunsuke/slog-util/slogutil"
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/urfave"
	"github.com/urfave/cli/v3"
)

func Run(ctx context.Context, logger *slogutil.Logger, env *urfave.Env) error {
	return urfave.Command(env, &cli.Command{ //nolint:wrapcheck
		Name:  "ghproj",
		Usage: "Add GitHub Issues and Pull Requests to GitHub Projects",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "log-level",
				Usage: "log level",
			},
		},
		Commands: []*cli.Command{
			(&initCommand{}).command(logger),
			(&addCommand{}).command(logger),
		},
	}).Run(ctx, env.Args)
}
