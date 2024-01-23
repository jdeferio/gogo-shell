package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		// start the prompt for gogo-shell
		fmt.Print("ggsh> ")

		command, args := collectInput()
		switch command {
		case "exit":
			os.Exit(0)
		case "":
			continue
		case "cd":
			cd(args)
		case "pwd":
			pwd()
		default:
			cmd := generateCmd(command, args)

			err := cmd.Run()
			if err != nil {
				log.Printf("Command finished with error: %v", err)
			}
		}
	}
}

func collectInput() (command string, args []string) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
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

func pwd() {
	dir, _ := os.Getwd()
	fmt.Println(dir)
}
