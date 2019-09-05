package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/wizcas/arse/loader"

	"github.com/urfave/cli"
)

const (
	defaultConfigFile = ".arse.yml"
)

func main() {
	app := cli.NewApp()
	app.Name = "arse"
	app.Usage = "Run custom actions with your own config"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Specify the config file instead of './.arse.yml'",
		},
	}
	app.Action = run
	app.Commands = []cli.Command{
		{
			Name:      "run",
			Aliases:   []string{"r"},
			Usage:     `Run a configured action by name`,
			UsageText: "arse [run|r] <ACTION_NAME>",
			Description: `This command finds the action specified by ACTION_NAME and execute it.
	 It's also the default command so you can omit the 'run' part.
	 For example, the following 2 commands:
		arse hello
		arse run hello
	 both run an action named 'hello' which is defined in the config file.`,
			Action: run,
		},
		{
			Name:     "init",
			Category: "Utilities",
			Usage:    "Create a new arse config file of chosen type with a bootstrap template",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "type, t",
					Usage: "The type of arse config file",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("Sorry but I haven't done it yet.")
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	configFile := c.String("config")
	if len(configFile) == 0 {
		configFile = defaultConfigFile
	}
	println("Hello my arse!")
	af, err := loader.SmartLoad(configFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Loaded ArseFile:\n%v\n", af)
	actionName := c.Args().Get(0)
	if len(actionName) == 0 {
		fmt.Println(`You need to specify an action to run.
Usage:
	arse [run|r] hello` + "\n")
		return errors.New("action not specified")
	}
	fmt.Printf("\naction to run: %s\n", actionName)

	action := af.Action(actionName)
	if action == nil {
		log.Fatalf("action not found: %s", actionName)
	}
	if err = action.Run(); err != nil {
		log.Fatal("action running error:\n%v", err)
	}
	return nil
}
