package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for true {
		fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			panic("Error reading command from the user.")
		}

		fmt.Println(command[:len(command)-1] + ": command not found")
		fmt.Println("hello")
	}
}
