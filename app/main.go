package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func exit(arguments []string) {
	if len(arguments) > 1 || len(arguments) == 0 {
		fmt.Println("Incorrect number of arguments")
	}

	code, err := strconv.Atoi(arguments[0])

	if err != nil {
		fmt.Println("Error converting string to int")
	}

	os.Exit(code)
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		inputString, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			os.Exit(1)
		}

		var input []string = strings.Split(inputString[:len(inputString)-1], " ")
		var command = input[0]
		var arguments = input[1:]

		switch command {
		case "exit":
			exit(arguments)
		case "echo":
			fmt.Println(strings.Join(arguments, " "))
		case "type":
			if arguments[0] == "echo" || arguments[0] == "exit" || arguments[0] == "type" {
				fmt.Println(arguments[0] + " is a shell builtin")
			} else {
				fmt.Println(arguments[0] + ": not found")
			}
		default:
			fmt.Println(command + ": command not found")
		}

	}
}
