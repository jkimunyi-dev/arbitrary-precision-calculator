package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Add Operation = iota
	Subtract
	Multiply
	Divide
	Pow
	Factorial
	Modulo
)

// REPL handles the Read-Eval-Print Loop for the calculator
type REPL struct{}

// Start begins the interactive calculator session
func (r *REPL) Start() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Arbitrary Precision Integer Calculator")
	fmt.Println("Enter expressions like: 123 + 456 or 10 ^ 5")
	fmt.Println("Type 'exit' or 'quit' to end the session")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading input:", err)
			continue
		}

		// Trim whitespace and newline
		input = strings.TrimSpace(input)

		// Check for exit commands
		if input == "exit" || input == "quit" {
			fmt.Println("Goodbye!")
			break
		}

		// Process the input
		result, err := r.processExpression(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("Result:", result)
	}
}
