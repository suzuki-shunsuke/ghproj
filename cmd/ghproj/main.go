package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/ghproj/pkg/cli"
	"github.com/suzuki-shunsuke/go-stdutil"
	"github.com/suzuki-shunsuke/logrus-error/logerr"
	"github.com/suzuki-shunsuke/logrus-util/log"
)

var (
	version = ""
	commit  = "" //nolint:gochecknoglobals
	date    = "" //nolint:gochecknoglobals
)

func main() {
	if code := run(); code != 0 {
		os.Exit(code)
	}
}

func run() int {
	logE := log.New("ghproj", version)
	if err := core(logE); err != nil {
		logerr.WithError(logE, err).Error("ghproj failed")
		return 1
	}
	return 0
}

func core(logE *logrus.Entry) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	return cli.Run(ctx, logE, &stdutil.LDFlags{ //nolint:wrapcheck
		Version: version,
		Commit:  commit,
		Date:    date,
	}, os.Args...)
}
