package app

type EditCommand struct {

}

var e EditCommand

func init()  {
	parser.AddCommand("e",
		"Edit information about a todo",
		"The edit command edits a todo. 'todo help e' to see more options.",
		&e)
}

/**
Usage:
e <id> <new_task>
e <id> -p <new_proj>
 */
func (cmd *EditCommand) Execute(args []string) error {
	return nil
}
