package main

import (
	"github.com/suzuki-shunsuke/ghproj/pkg/cli"
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/urfave"
)

var version = ""

func main() {
	urfave.Main("ghproj", version, cli.Run)
}
