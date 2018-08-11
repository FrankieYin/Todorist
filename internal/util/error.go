package util

import (
	"fmt"
	"os"
)

type InvalidIdError struct {
	Msg string
}

type ProjectNotFound struct {
	Msg string
}

type TooManyArguments struct {
	Msg string
}

type NotEnoughArguments struct {
	Msg string
}

func (e InvalidIdError) Error() string {
	return e.Msg
}

func (e ProjectNotFound) Error() string {
	return e.Msg
}

func (e TooManyArguments) Error() string {
	return e.Msg
}

func (e NotEnoughArguments) Error() string {
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
