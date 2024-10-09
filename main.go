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
	cfg := configuration{
		previous: "",
		next:     "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
	}
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
			err = command.callback(&cfg)
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
