package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for 0 < 1 {
		// 	//start the prompt for gogo-shell
		fmt.Print("ggsh> ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "exit" {
			break
		}

		cmd := exec.Command(input)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		cmd.Run()
	}
}

// func execute_cmd(input string) {
// 	cmd := exec.Command(input)
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr

// 	cmd.Run()
// }

// func collect_input() {
// 	//start the prompt for gogo-shell
// 	fmt.Print("ggsh> ")

// 	reader := bufio.NewReader(os.Stdin)
// 	input, _ := reader.ReadString('\n')
// 	input = strings.TrimSpace(input)
// 	return input
// }
