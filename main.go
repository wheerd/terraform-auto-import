package main

import (
	"log"
	"os"

	"github.com/wheerd/terraform-auto-import/v2/core"

	"github.com/urfave/cli/v2"
)

func main() {
	var runconfig core.RunConfig

	app := &cli.App{
		Name:  "terraform-auto-import",
		Usage: "Automatically import existing AWS resources into terraform",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "tfplan",
				Usage:       "Load terraform plan from `FILE`",
				Destination: &runconfig.TerraformPlanPath,
			},
		},
		Action: func(c *cli.Context) error {
			return core.Run(&runconfig)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
