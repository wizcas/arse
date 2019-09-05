package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wizcas/arse/loaders/yaml"

	"github.com/urfave/cli"
)

const (
	defaultConfigFile = ".arse.yml"
)

func main() {
	app := cli.NewApp()
	app.Name = "arse"
	app.Usage = "Run custom actions with your own config"
	app.Commands = []cli.Command{
		{
			Name:      "run",
			ShortName: "r",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "config, c",
					Usage: "Specify the config file instead of './.arse.yml'",
				},
			},
			Action: func(c *cli.Context) error {
				configFile := c.String("config")
				if len(configFile) == 0 {
					configFile = defaultConfigFile
				}
				run(configFile)
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(filename string) {
	println("Hello my arse!")
	data, err := yaml.Load(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Loaded ArseFile:\n%v\n", data)
}
