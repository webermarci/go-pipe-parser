package pipeparser

import (
	"bytes"
	"os/exec"
	"strings"

	pipe "github.com/b4b4r07/go-pipe"
)

const (
	stateStart  = "start"
	stateArgs   = "args"
	stateQuotes = "quotes"

	singleQuoteChar = '\''
	doubleQuoteChar = '"'
	backslashChar   = '\\'
	tabChar         = '\t'
	spaceChar       = ' '
)

func parse(input string) []string {
	var args []string
	state := stateStart
	current := ""
	quote := string(doubleQuoteChar)
	isEscapeNext := true

	for i := 0; i < len(input); i++ {
		c := input[i]

		if state == stateQuotes {
			if string(c) != quote {
				current += string(c)
			} else {
				args = append(args, current)
				current = ""
				state = stateStart
			}
			continue
		}

		if isEscapeNext {
			current += string(c)
			isEscapeNext = false
			continue
		}

		if c == backslashChar {
			isEscapeNext = true
			continue
		}

		if c == doubleQuoteChar || c == singleQuoteChar {
			state = stateQuotes
			quote = string(c)
			continue
		}

		if state == stateArgs {
			if c == spaceChar || c == tabChar {
				args = append(args, current)
				current = ""
				state = stateStart
			} else {
				current += string(c)
			}
			continue
		}

		if c != spaceChar && c != tabChar {
			state = stateArgs
			current += string(c)
		}
	}

	if current != "" {
		args = append(args, current)
	}

	return args
}

func buildCommands(input string) []*exec.Cmd {
	commands := []*exec.Cmd{}
	commandStrings := strings.Split(input, " | ")
	for _, s := range commandStrings {
		parsed := parse(s)
		command := exec.Command(parsed[0], parsed[1:]...)
		commands = append(commands, command)
	}
	return commands
}

// Run executes the pipe command
func Run(input string) (bytes.Buffer, error) {
	var b bytes.Buffer
	pipeErr := pipe.Command(&b, buildCommands(input)...)
	if pipeErr != nil {
		return b, pipeErr
	}
	return b, nil
}
