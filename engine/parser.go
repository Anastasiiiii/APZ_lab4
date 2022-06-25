package engine

import (
	"strconv"
	"strings"
)

func Parse(commandLine string) Command {
	cmdFields := strings.Fields(commandLine)
	availableCommands := "Available commands: print (1 argument), add (2 arguments)."

	if len(cmdFields) < 2 {
		return &ErrorMessage{message: "Incorrect number of args or no command. " + availableCommands}
	}

	command := cmdFields[0]

	switch command {
	case "print":
		if (len(cmdFields)) != 2 {
			return &ErrorMessage{message: "`print` command can accept only one argument"}
		}

		return &PrintCommand{arg: cmdFields[1]}

	case "add":
		if len(cmdFields) != 3 {
			return &ErrorMessage{message: "`add` command requires exactly 2 arguments"}
		}

		arg1, err := strconv.Atoi(cmdFields[1])
		if err != nil {
			return &ErrorMessage{message: err.Error()}
		}

		arg2, err := strconv.Atoi(cmdFields[2])
		if err != nil {
			return &ErrorMessage{message: err.Error()}
		}

		return &AddCommand{arg1: arg1, arg2: arg2}

	default:
		return &ErrorMessage{message: "Unknown instruction. " + availableCommands}
	}
}
