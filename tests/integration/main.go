package main

import (
	"fmt"
	"log"
	"os"

	authTests "github.com/alexandr-io/backend/auth/tests"
	libraryTests "github.com/alexandr-io/backend/library/tests"
	userTests "github.com/alexandr-io/backend/user/tests"
	"github.com/urfave/cli/v2"
)

var authToken string
var asynchronous = false

var flags = []cli.Flag{
	&cli.StringSliceFlag{
		Name:    "include",
		Aliases: []string{"i"},
		Usage:   "include tests for `TEST_SUIT`",
	},
	&cli.StringSliceFlag{
		Name:    "exclude",
		Aliases: []string{"e"},
		Usage:   "exclude tests for `TEST_SUIT`",
	},
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "asynchronous",
				Aliases: []string{"a"},
				Usage:   "run the tests asynchronously",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "tests",
				Usage: "print all test suit names",
				Action: func(_ *cli.Context) error {
					printTestSuits()
					return nil
				},
			},
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
	}
	app.UseShortOptionHandling = true
	app.EnableBashCompletion = true

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func parseAndExecTests(c *cli.Context, environment string) error {
	var errorHappened = false
	testsToExec := filters(c)

	if err := execBasicAuth(environment); err != nil {
		return cli.Exit("", 1)
	}

	for _, value := range testsToExec {
		value := value
		if asynchronous {
			go func() {
				if err := value.Func(environment); err != nil {
					errorHappened = true
				}
				value.Channel <- true
			}()
		} else {
			if err := value.Func(environment); err != nil {
				errorHappened = true
			}
		}
	}

	// Wait for goroutines to finish is test are run asynchronously
	if asynchronous {
		for _, value := range testsToExec {
			<-value.Channel
		}
	}

	if errorHappened {
		return cli.Exit("", 1)
	}
	return nil
}

func printTestSuits() {
	fmt.Println("USAGE:\n   Using the MICROSERVICE_NAME will include or exclude all of the TEST_SUITS\n")
	fmt.Println("FORMAT:\n   • MICROSERVICE_NAME\n     · TEST_SUITS\n")
	fmt.Println("TEST SUITS LIST:")
	for _, service := range testSuits {
		fmt.Printf("   • %s\n", service.Microservice)
		for _, suit := range service.Suits {
			fmt.Printf("     · %s\n", suit.Name)
		}
	}
}

func execLibrary(environment string) error {
	if err := libraryTests.ExecLibraryTests(environment, authToken); err != nil {
		return err
	}
	return nil
}

func execUser(environment string) error {
	if err := userTests.ExecUserTests(environment, authToken); err != nil {
		return err
	}
	return nil
}

func execBasicAuth(environment string) error {
	jwt, err := authTests.ExecAuthBasicTests(environment)
	if err != nil {
		return err
	}
	authToken = jwt
	return nil
}
