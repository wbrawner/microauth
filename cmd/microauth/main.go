package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/wbrawner/microauth/internal/microauth"
)

const FLAG_API_KEY string = "api-key"
const FLAG_PORT string = "port"

func main() {
	app := &cli.App{
		Name:  "microauth",
		Usage: "a microservice for authenticating users",
		// TODO: Add flags for database connection parameters
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    FLAG_API_KEY,
				Aliases: []string{"a"},
				Usage:   "`KEY` to use when authenticating with the server",
				EnvVars: []string{
					"MICROAUTH_API_KEY",
				},
			},
			&cli.IntFlag{
				Name:    FLAG_PORT,
				Aliases: []string{"p"},
				Value:   8080,
				Usage:   "`PORT` for server to listen on",
				EnvVars: []string{
					"MICROAUTH_PORT",
				},
			},
		},
		Action: func(c *cli.Context) error {
			return microauth.StartServer(c.Int(FLAG_PORT), c.String(FLAG_API_KEY))
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
