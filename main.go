package main

import (
	"log"
	"os"

	"github.com/wheerd/terraform-auto-import/v2/commands"

	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("terraform-auto-import", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"list-new": func() (cli.Command, error) {
			return &commands.ListNewResourcesCommand{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
