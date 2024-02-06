package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		// start the prompt for gogo-shell
		fmt.Print("ggsh> ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		inputSplit := SplitPipes(input)

		var cmds []*exec.Cmd
		var output io.ReadCloser

		for _, rawInput := range inputSplit {
			command, args := ParseInput(rawInput)

			switch command {
			case "exit":
				os.Exit(0)
			case "":
				continue
			case "cd":
				cd(args)
			case "pwd":
				pr, pw := io.Pipe()
				go pwd(pw)
			default:
				cmd := generateCmd(command, args)

				err := cmd.Run()
				if err != nil {
					log.Printf("Command finished with error: %v", err)
				}
			}

		}

	}
}

func SplitPipes(input string) (cmds []string) {
	cmds = strings.Split(input, "|")

	return cmds
}

func ParseInput(input string) (command string, args []string) {
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)

	if len(parts) >= 1 {
		return parts[0], parts[1:]
	}

	return "", []string{}
}

func generateCmd(command string, args []string) (cmd *exec.Cmd) {
	cmd = exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}

func cd(args []string) {
	var path string

	var err error

	if len(args) > 0 {
		path = args[0]
	} else {
		path, _ = os.UserHomeDir()
	}

	err = os.Chdir(path)
	if err != nil {
		log.Printf("%v\n", err)
	}
}

func pwd(w *io.PipeWriter) {
	defer w.Close()
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(w, dir)
}
