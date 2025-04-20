package cli

import (
	"context"

	"github.com/suzuki-shunsuke/urfave-cli-v3-util/urfave"
	"github.com/urfave/cli/v3"
)

func Run(ctx context.Context, runner *urfave.Runner, args ...string) error {
	return runner.Command(&cli.Command{ //nolint:wrapcheck
		Name:    "ghproj",
		Usage:   "",
		Version: runner.LDFlags.Version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "log-level",
				Usage: "log level",
			},
			&cli.StringFlag{
				Name:  "log-color",
				Usage: "Log color. One of 'auto' (default), 'always', 'never'",
			},
		},
		EnableShellCompletion: true,
		Commands: []*cli.Command{
			(&initCommand{
				logE: runner.LogE,
			}).command(),
			(&addCommand{
				logE: runner.LogE,
			}).command(),
			(&completionCommand{
				logE:   runner.LogE,
				stdout: runner.Stdout,
			}).command(),
		},
	}).Run(ctx, args)
}
