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
		// 	//start the prompt for gogo-shell
		fmt.Print("ggsh> ")

		commandName, args := collectInput()
		if commandName == "exit" {
			os.Exit(0)
		} else if commandName == "" {
			continue
		}

		cmd := generateCmd(commandName, args)

		err := cmd.Run()
		if err != nil {
			log.Printf("Command finished with error: %v", err)
		}
	}
}

func collectInput() (commandName string, args []string) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")

	return parts[0], parts[1:]
}

func generateCmd(commandName string, args []string) (cmd *exec.Cmd) {
	cmd = exec.Command(commandName, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}
