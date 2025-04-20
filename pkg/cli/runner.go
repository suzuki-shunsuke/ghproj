package cli

import (
	"context"

	"github.com/suzuki-shunsuke/urfave-cli-v3-util/urfave"
	"github.com/urfave/cli/v3"
)

type Runner struct {
	Runner *urfave.Runner
}

func (r *Runner) Run(ctx context.Context, args ...string) error {
	return r.Runner.Command(&cli.Command{ //nolint:wrapcheck
		Name:    "ghproj",
		Usage:   "",
		Version: r.Runner.LDFlags.Version,
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
				logE: r.Runner.LogE,
			}).command(),
			(&addCommand{
				logE: r.Runner.LogE,
			}).command(),
			(&completionCommand{
				logE:   r.Runner.LogE,
				stdout: r.Runner.Stdout,
			}).command(),
		},
	}).Run(ctx, args)
}
