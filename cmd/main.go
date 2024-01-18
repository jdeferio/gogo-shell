package main

import ()

// Print the shell prompt
// Accept user input
// preprocess user input (trim white space; detect delimiter)
// delim == \n
// execute input command
// terminate shell after executing

func main() {
	//start the prompt for gogo-shell
	fmt.Print("ggsh> ")

	//
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

}

func preprocess() {

}
