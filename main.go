package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Scanner is used for input handling
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		printReplLine()
		scanner.Scan()
		fmt.Println()
		err := scanner.Err()
		if err != nil {
			fmt.Println("Error reading command\n", err)
			continue
		}
		// Lower case and remove trailing/leading whitespace to improve command recognition
		userCommand := strings.ToLower(strings.TrimSpace(scanner.Text()))
		command, ok := commands[userCommand]
		if ok {
			err = command.callback()
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
			}
		} else {
			fmt.Printf("Command '%s' doesn't exist\n", userCommand)
		}
		fmt.Println()
	}
}

func printReplLine() {
	fmt.Print("Pokedex > ")
}
