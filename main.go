package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
	"sps-recruitment/application"
)

func main() {

	clientApp := cli.NewApp()
	clientApp.Name = "sps-recruitment"
	clientApp.Version = "0.0.1"
	clientApp.Commands = []cli.Command{
		{
			Name:        "sum",
			Description: "start sum",
			Action: func(c *cli.Context) error {
				application.GetSumResult()
				return nil
			},
		},
		{
			Name:        "contains",
			Description: "start contains",
			Action: func(c *cli.Context) error {
				application.GetContainsResult()
				return nil
			},
		},
		{
			Name:        "person",
			Description: "new person",
			Action: func(c *cli.Context) error {
				application.GetPerson()
				return nil
			},
		},
		{
			Name:        "rearrange",
			Description: "start unit test",
			Action: func(c *cli.Context) error {
				application.GetRearrange()
				return nil
			},
		},
	}

	err := clientApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}