package cli

import (
	"net/http"

	"github.com/m-mizutani/goerr"
	"github.com/m-mizutani/httpdump/pkg/server"
	"github.com/m-mizutani/httpdump/pkg/utils"
	"github.com/urfave/cli/v2"
)

func runServe() *cli.Command {
	var (
		addr string
	)

	return &cli.Command{
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "Run HTTP dump server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "addr",
				Aliases:     []string{"a"},
				Value:       "0.0.0.0:8080",
				Usage:       "Listen address",
				EnvVars:     []string{"HTTPDUMP_ADDR"},
				Destination: &addr,
			},
		},
		Action: func(ctx *cli.Context) error {
			srv := server.New()

			utils.Logger().Info("Start HTTP dump server")
			if err := http.ListenAndServe(addr, srv); err != nil {
				return goerr.Wrap(err, "failed to listen and serve")
			}

			return nil
		},
	}
}
