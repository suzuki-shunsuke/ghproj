package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/ghproj/pkg/cli"
	"github.com/suzuki-shunsuke/logrus-error/logerr"
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/log"
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/urfave"
)

var (
	version = ""
	commit  = "" //nolint:gochecknoglobals
	date    = "" //nolint:gochecknoglobals
)

func main() {
	logE := log.New("ghproj", version)
	if err := core(logE); err != nil {
		logerr.WithError(logE, err).Fatal("ghproj failed")
	}
}

func core(logE *logrus.Entry) error {
	runner := &cli.Runner{
		Runner: &urfave.Runner{
			Stdin:  os.Stdin,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
			LDFlags: &urfave.LDFlags{
				Version: version,
				Commit:  commit,
				Date:    date,
			},
			LogE: logE,
		},
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	return runner.Run(ctx, os.Args...) //nolint:wrapcheck
}
