package main

import (
	"fmt"
	"log"
	"os"

	"arbitrary-precision-calculator/internal"
)

type REPL struct{}

func (r *REPL) Start() {
	// Implement REPL logic here
	fmt.Println("Starting REPL...")
}

func main() {
	// If no arguments, start REPL
	if len(os.Args) == 1 {
		repl := &REPL{}
		repl.Start()
		return
	}

	// Rest of your main function remains the same
	// If arguments are provided, attempt to parse as a calculation
	if len(os.Args) < 4 {
		fmt.Println("Usage: calculator [num1] [operation] [num2]")
		fmt.Println("Or run without arguments to start interactive REPL")
		os.Exit(1)
	}

	// Parse first number
	num1, err := internal.NewArbitraryInt(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid first number: %v", err)
	}

	// Parse operation
	operation := os.Args[2]

	// Parse second number
	num2, err := internal.NewArbitraryInt(os.Args[3])
	if err != nil {
		log.Fatalf("Invalid second number: %v", err)
	}

	// Perform calculation based on operation
	switch operation {
	case "+":
		fmt.Println(num1.Add(num2))
	case "-":
		fmt.Println(num1.Subtract(num2))
	case "*":
		fmt.Println(num1.Multiply(num2))
	case "/":
		quotient, remainder, err := num1.Divide(num2)
		if err != nil {
			log.Fatalf("Division error: %v", err)
		}
		fmt.Printf("Quotient: %v\nRemainder: %v\n", quotient, remainder)
	case "^":
		result, err := num1.Pow(num2)
		if err != nil {
			log.Fatalf("Exponentiation error: %v", err)
		}
		fmt.Println(result)
	case "%":
		result, err := num1.Modulo(num2)
		if err != nil {
			log.Fatalf("Modulo error: %v", err)
		}
		fmt.Println(result)
	default:
		log.Fatalf("Unsupported operation: %s", operation)
	}
}
