package main

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/wizcas/arse/parser"
	"log"
	"os"
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
				fmt.Printf("running with file: %s\n", configFile)
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
	data, err := parser.Load(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Loaded ArseFile:\n%v\n", data)
}
