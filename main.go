package main

import (
	"log/slog"
	"os"

	"github.com/m-mizutani/httpdump/pkg/cli"
	"github.com/m-mizutani/httpdump/pkg/utils"
)

func main() {
	if err := cli.Run(os.Args); err != nil {
		utils.Logger().Error("Failed to run app", slog.Any("err", err))
	}
}
