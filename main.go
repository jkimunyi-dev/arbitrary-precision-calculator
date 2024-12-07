package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"arbitrary-precision-calculator/internal"
)

const (
	VERSION   = "1.0.0"
	HELP_TEXT = `Arbitrary Precision Calculator

Usage:
  arbitrary-precision-calculator [options] [num1] [operation] [num2]

Options:
  --help, -h      Show this help message
  --version, -v   Show version information

Operations:
  +               Addition
  -               Subtraction
  *               Multiplication
  /               Division (shows quotient and remainder)
  ^               Exponentiation
  %               Modulo

Examples:
  # Basic arithmetic
  arbitrary-precision-calculator 1000000000 + 2000000000
  arbitrary-precision-calculator 5 ^ 3

  # Different number formats
  arbitrary-precision-calculator 0b1010 + 0b1100      # Binary
  arbitrary-precision-calculator 0xff + 0x100         # Hexadecimal

  # Interactive REPL
  arbitrary-precision-calculator               # Start interactive mode

Build Information:
  Version:        %s
  Go Version:     %s
  Platform:       %s/%s
`
)

type REPL struct{}

func (r *REPL) Start() {
	fmt.Println("Arbitrary Precision Calculator REPL")
	fmt.Println("Type 'exit' or 'quit' to leave the REPL")
	fmt.Println("Enter calculations like: 1000 + 2000")
	// TODO: Implement full REPL logic
}

func printHelp() {
	fmt.Printf(HELP_TEXT,
		VERSION,
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
	)
}

func main() {
	// Handle help and version flags
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--help", "-h":
			printHelp()
			return
		case "--version", "-v":
			fmt.Printf("Arbitrary Precision Calculator v%s\n", VERSION)
			fmt.Printf("Build with %s for %s/%s\n",
				runtime.Version(),
				runtime.GOOS,
				runtime.GOARCH,
			)
			return
		}
	}

	// If no arguments, start REPL
	if len(os.Args) == 1 {
		repl := &REPL{}
		repl.Start()
		return
	}

	// Normalize operations to match internal methods
	operationMap := map[string]string{
		"add":  "+",
		"sub":  "-",
		"mult": "*",
		"div":  "/",
		"pow":  "^",
		"mod":  "%",
	}

	// If first argument matches an operation name, shift arguments
	if len(os.Args) == 3 && operationMap[os.Args[1]] != "" {
		os.Args = append([]string{os.Args[0], os.Args[2], operationMap[os.Args[1]], os.Args[2]})
	}

	// Validate argument count
	if len(os.Args) < 4 {
		fmt.Println("Usage: calculator [num1] [operation] [num2]")
		fmt.Println("Run with --help for more information")
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
		log.Fatalf("Unsupported operation: %s. Run with --help for supported operations.", operation)
	}
}
