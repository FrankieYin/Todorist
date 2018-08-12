package app

import (
	"fmt"
	"os"
	"github.com/FrankieYin/todo/internal/util"
	"github.com/FrankieYin/todo/internal/data"
)

type ProjCommand struct {
	Rename bool `short:"m" description:"Rename a project from <oldname> to <newname>."`
	Verbose []bool `short:"v" long:"verbose"`
	Note bool `short:"n" description:"Add a description to the project."`
	Delete bool `short:"d" description:"Delete a project specified by the <name>."`
}

var proj ProjCommand

func init() {
	parser.AddCommand("proj",
		"List, create, rename projects.",
		"List, create, rename projects.",
		&proj)
}

func (cmd *ProjCommand) Execute(args []string) error {
	n := len(args)
	if n == 0 { // list existing projects
		if len(projList.Projects) == 1 {
			fmt.Println("No existing project found.")
			fmt.Println("Use 'todo proj [-n <description>] <name>' to create a new project.")
			os.Exit(0)
		}

		for _, proj := range projList.Projects[1:] { // skip Inbox default project
			asterisk := " "
			if proj.OnFocus {
				asterisk = "*"
			}
			fmt.Printf("%s%s\n", asterisk, proj.Name)
		}
		os.Exit(0)
	}

	if proj.Delete { // usage: proj -d <project_name>
		if n > 1 {return util.TooManyArguments{Msg:"fatal: too many arguments for a delete operation"}}
		projList.DeleteProject(args[0])
		return save(projList, projJsonFilename)
	}

	if proj.Rename { // usage: proj -m <old_name> <new_name>
		if n < 2 {return util.NotEnoughArguments{Msg:"fatal: rename operation needs 2 arguments, 1 given"}}
		if n > 2 {return util.TooManyArguments{Msg:"fatal: too many arguments for a rename operation"}}
		projList.RenameProject(args[0], args[1])
		return save(projList, projJsonFilename)
	}

	// create a new project
	var p *data.Project
	if proj.Note { // usage: proj <name> [-n <description>]
		if n < 2 {return util.NotEnoughArguments{Msg:"fatal: proj -n operation needs 2 arguments, 1 given"}}
		if n > 2 {return util.TooManyArguments{Msg:"fatal: too many arguments for creating a project\nDid you enclose the description in a \"\" ?"}}
		p = &data.Project{Name:args[1], Description:args[0]}
	} else {
		p = &data.Project{Name:args[0]}
	}
	projList.AddProject(p)
	return save(projList, projJsonFilename)
}
