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
	SetFocus bool `long:"set-focus"`
	Focus bool `short:"f" long:"focus"`
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
		if len(data.ProjList.Projects) == 0 {
			fmt.Println("No existing project found.")
			fmt.Println("Use 'todo proj [-n <description>] <name>' to create a new project.")
			os.Exit(0)
		}

		for _, proj := range data.ProjList.Projects { // skip Inbox default project
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
		data.ProjList.DeleteProject(args[0])
	} else if proj.Rename { // usage: proj -m <old_name> <new_name>
		if n < 2 {return util.NotEnoughArguments{Msg:"fatal: rename operation needs 2 arguments, 1 given"}}
		if n > 2 {return util.TooManyArguments{Msg:"fatal: too many arguments for a rename operation"}}
		data.ProjList.RenameProject(args[0], args[1])
	} else if proj.Focus { // proj -f <project_name>
		// flip the onFocus state of the project
		if err := data.ProjList.ChangeFocus(args); err != nil {return err}
	} else { // create a new project
		var p *data.Project
		if proj.Note { // usage: proj [--set-focus] <name> [-n <description>]
			if n < 2 {return util.NotEnoughArguments{Msg:"fatal: proj -n operation needs 2 arguments, 1 given"}}
			if n > 2 {return util.TooManyArguments{Msg:"fatal: too many arguments for creating a project\nDid you enclose the description in a \"\" ?"}}
			p = &data.Project{Name:args[0], Description:args[1], OnFocus:proj.SetFocus}
		} else {
			p = &data.Project{Name:args[0], OnFocus:proj.SetFocus}
		}
		data.ProjList.AddProject(p)
	}

	return save()
}
