package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IncorrectArgumentsSupplied() {
	fmt.Println("Incorrect number of arguments")
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

		if command == "exit" {
			if len(arguments) > 1 || len(arguments) == 0 {
				IncorrectArgumentsSupplied()
			}
			fmt.Println()

			code, err := strconv.Atoi(arguments[0])

			if err != nil {
				fmt.Println("Error converting string to int")
			}

			os.Exit(code)
		}

		fmt.Println(command + ": command not found")
	}
}
