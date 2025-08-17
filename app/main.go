package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var buildinCommands = map[string]func([]string){
	"exit": handleExit,
	"echo": handleEcho,
	"pwd":  handlePWD,
}

func handleExit(arguments []string) {
	if len(arguments) > 1 {
		fmt.Println("Incorrect number of arguments")
		return
	}

	if len(arguments) == 0 {
		os.Exit(0)
	}

	code, err := strconv.Atoi(arguments[0])

	if err != nil {
		fmt.Println("Error converting string to int")
		return
	}

	os.Exit(code)
}

func handleEcho(arguments []string) {
	fmt.Println(strings.Join(arguments, " "))
}

func isFileExecutable(filePath string) (string, bool) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return filePath, false
	}

	mode := fileInfo.Mode()
	return filePath, mode.Perm()&0100 != 0
}

func findExecutableInPath(command string) {
	paths := strings.SplitSeq(os.Getenv("PATH"), ":")
	for path := range paths {
		dir, err := os.Open(path)

		if err != nil {
			continue
		}

		files, err := dir.ReadDir(-1)
		if err != nil {
			continue
		}

		for _, file := range files {
			if file.Name() == command {
				filePath := dir.Name() + "/" + file.Name()
				if _, isIt := isFileExecutable(filePath); isIt {
					fmt.Println(command + " is " + filePath)
					return
				}
			}
		}
	}

	fmt.Println(command + ": not found")
}

func handleType(arguments []string) {
	for _, argument := range arguments {
		_, ok := buildinCommands[argument]
		if argument != "type" && !ok {
			findExecutableInPath(argument)
			continue
		}
		fmt.Println(argument + " is a shell builtin")
	}
}

func handlePWD(arguments []string) {
	path, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Println(path)
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

		handler, ok := buildinCommands[command]
		if ok {
			handler(arguments)
			continue
		}

		if command == "type" {
			handleType(arguments)
			continue
		}

		cmd := exec.Command(command, arguments...)
		var out strings.Builder
		cmd.Stdout = &out
		if err := cmd.Run(); err != nil {
			fmt.Println(command + ": command not found")
			continue
		}
		fmt.Print(out.String())
	}
}
