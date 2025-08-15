package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	command, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		panic("Error reading command from the user.")
	}

	// take all but the last character
	// because the ReadString function returns the
	// whole read string upto and including the delimiter
	fmt.Println(command[:len(command)-1] + ": command not found")
}
