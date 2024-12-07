package main

import (
	"arbitrary-precision-calculator/internal"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

// processExpression parses and evaluates a mathematical expression
func (r *REPL) processExpression(input string) (string, error) {
	// Remove extra whitespace
	input = regexp.MustCompile(`\s+`).ReplaceAllString(input, " ")

	// Split the input into parts
	parts := strings.Split(input, " ")

	// Handle special case for factorial
	if len(parts) == 2 && (parts[1] == "!" || parts[1] == "factorial") {
		return r.handleFactorial(parts[0])
	}

	// Ensure we have a valid expression (num op num)
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid expression format")
	}

	// Parse first number
	num1, err := internal.NewArbitraryInt(parts[0])
	if err != nil {
		return "", fmt.Errorf("invalid first number: %v", err)
	}

	// Parse operation
	op, err := r.parseOperation(parts[1])
	if err != nil {
		return "", err
	}

	// Parse second number
	num2, err := internal.NewArbitraryInt(parts[2])
	if err != nil {
		return "", fmt.Errorf("invalid second number: %v", err)
	}

	// Perform the operation
	return r.performOperation(num1, num2, op)
}

// parseOperation converts a string to an Operation
func (r *REPL) parseOperation(opStr string) (Operation, error) {
	switch strings.ToLower(opStr) {
	case "+":
		return Add, nil
	case "-":
		return Subtract, nil
	case "*":
		return Multiply, nil
	case "/":
		return Divide, nil
	case "^":
		return Pow, nil
	case "%":
		return Modulo, nil
	default:
		return Add, fmt.Errorf("unsupported operation: %s", opStr)
	}
}

// performOperation executes the specified operation
func (r *REPL) performOperation(num1, num2 *internal.ArbitraryInt, op Operation) (string, error) {
	switch op {
	case Add:
		return num1.Add(num2).String(), nil
	case Subtract:
		return num1.Subtract(num2).String(), nil
	case Multiply:
		return num1.Multiply(num2).String(), nil
	case Divide:
		quotient, _, err := num1.Divide(num2)
		if err != nil {
			return "", err
		}
		return quotient.String(), nil
	case Pow:
		result, err := num1.Pow(num2)
		if err != nil {
			return "", err
		}
		return result.String(), nil
	case Modulo:
		result, err := num1.Modulo(num2)
		if err != nil {
			return "", err
		}
		return result.String(), nil
	default:
		return "", fmt.Errorf("unsupported operation")
	}
}

// handleFactorial performs factorial operation
func (r *REPL) handleFactorial(numStr string) (string, error) {
	num, err := internal.NewArbitraryInt(numStr)
	if err != nil {
		return "", fmt.Errorf("invalid number for factorial: %v", err)
	}

	result, err := num.Factorial()
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
