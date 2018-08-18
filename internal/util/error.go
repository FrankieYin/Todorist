package util

import (
	"fmt"
	"os"
)

type InvalidIdError struct {
	Msg string
}

type ProjectNotFound struct {
	Name string
}

type TooManyArguments struct {
	Msg string
}

type NotEnoughArguments struct {
	Msg string
}

type InvalidArgument struct {
	Msg string
}

type InvalidPriorityLevel struct {
	Level int
}

func (e InvalidPriorityLevel) Error() string {
	msg := fmt.Sprintf("Does not have a priority level %d.\n " +
		"A general rule of thumbs:\n" +
		"Use 4 for non-urgent and non-important tasks;\n" +
		"Use 3 for non-urgent and important tasks;\n" +
		"Use 2 for urgent and non-important tasks;\n" +
		"Use 1 for urgent and important tasks.", e.Level)
	return msg
}

func (e InvalidIdError) Error() string {
	return e.Msg
}

func (e ProjectNotFound) Error() string {
	msg := fmt.Sprintf("Project %s does not exist.\n", e.Name)
	return msg
}

func (e TooManyArguments) Error() string {
	return e.Msg
}

func (e NotEnoughArguments) Error() string {
	return e.Msg
}

func (e InvalidArgument) Error() string {
	return e.Msg
}

func CheckErr(err error, msg string) {
	if err != nil {
		if msg != "" {
			fmt.Println(msg)
		}
		fmt.Println(err)
		os.Exit(1)
	}
}
