package command

import (
	"fmt"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/dnaeon/gru/module"
	"github.com/gosuri/uitable"
)

// NewModuleCommand creates a new sub-command for
// displaying the list of available modules
func NewModuleCommand() cli.Command {
	cmd := cli.Command{
		Name:   "module",
		Usage:  "display available modules",
		Action: execModuleCommand,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "siterepo",
				Value:  "",
				Usage:  "path/url to the site repo",
				EnvVar: "GRU_SITEREPO",
			},
		},
	}

	return cmd
}

// Executes the "module" command
func execModuleCommand(c *cli.Context) {
	modulePath := filepath.Join(c.String("siterepo"), "modules")
	registry, err := module.Discover(modulePath)
	if err != nil {
		displayError(err, 1)
	}

	table := uitable.New()
	table.AddRow("MODULE", "PATH")

	for m, p := range registry {
		table.AddRow(m, p)
	}

	fmt.Println(table)
}
