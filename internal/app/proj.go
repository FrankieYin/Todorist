package app

import (
	"fmt"
	"os"
	)

type ProjCommand struct {
}

var proj ProjCommand

func init() {
	parser.AddCommand("proj",
		"List, create, rename projects.",
		"List, create, rename projects.",
		&proj)
}

func (cmd *ProjCommand) Execute(args []string) error {

	projList = loadProject(projJsonFilename)

	n := len(args)
	if n == 0 { // list existing projects
		for _, proj := range projList.Projects {
			asterisk := " "
			if proj.OnFocus {
				asterisk = "*"
			}
			fmt.Printf("%s%s\n", asterisk, proj.Name)
		}
		os.Exit(0)
	}

	// parse the flags
	return nil
}
