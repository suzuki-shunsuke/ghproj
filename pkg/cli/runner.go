package cli

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-stdutil"
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/urfave"
	"github.com/urfave/cli/v3"
)

func Run(ctx context.Context, logE *logrus.Entry, ldFlags *stdutil.LDFlags, args ...string) error {
	return urfave.Command(ldFlags, &cli.Command{ //nolint:wrapcheck
		Name:    "ghproj",
		Usage:   "",
		Version: ldFlags.Version,
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
				logE: logE,
			}).command(),
			(&addCommand{
				logE: logE,
			}).command(),
			(&completionCommand{
				logE:   logE,
				stdout: os.Stdout,
			}).command(),
		},
	}).Run(ctx, args)
}
