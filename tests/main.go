package main

import (
	"log"
	"os"

	authTests "github.com/alexandr-io/backend/auth/tests"
	"github.com/urfave/cli/v2"
)

var flags = []cli.Flag{
	&cli.StringSliceFlag{
		Name:    "include",
		Aliases: []string{"i"},
		Usage:   "Include tests for `MICRO_SERVICE`",
	},
	&cli.StringSliceFlag{
		Name:    "exclude",
		Aliases: []string{"e"},
		Usage:   "Exclude tests for `MICRO_SERVICE`",
	},
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "test",
				Usage: "options for running tests",
				Subcommands: []*cli.Command{
					{
						Name:  "local",
						Usage: "run tests in local",
						Flags: flags,
						Action: func(c *cli.Context) error {
							return parseAndExecTests(c, "local")
						},
					},
					{
						Name:  "preprod",
						Usage: "run tests in preprod",
						Flags: flags,
						Action: func(c *cli.Context) error {
							return parseAndExecTests(c, "preprod")
						},
					},
					{
						Name:  "prod",
						Usage: "run tests in prod",
						Flags: flags,
						Action: func(c *cli.Context) error {
							return parseAndExecTests(c, "prod")
						},
					},
				},
			},
		},
	}
	app.EnableBashCompletion = true

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func parseAndExecTests(c *cli.Context, environment string) error {
	testsToExec, err := filters(c)
	if err != nil {
		return err
	}
	for _, elem := range testsToExec {
		if err := elem(environment); err != nil {
			return cli.Exit("", 1)
		}
	}
	return nil
}

func execUser(environment string) error {
	return nil
}

func execAuth(environment string) error {
	if err := authTests.ExecAuthTests(environment); err != nil {
		return err
	}
	return nil
}
