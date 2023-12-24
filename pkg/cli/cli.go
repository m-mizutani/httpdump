package cli

import (
	"github.com/m-mizutani/goerr"
	"github.com/urfave/cli/v2"
)

func Run(args []string) error {
	app := cli.App{
		Name:  "httpdump",
		Usage: "HTTP dump server",
		Commands: []*cli.Command{
			runServe(),
		},
	}

	if err := app.Run(args); err != nil {
		return goerr.Wrap(err, "failed to run app")
	}

	return nil
}
