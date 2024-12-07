package internal

import (
	"fmt"
	"strings"
)

// ArbitraryInt represents an arbitrary precision integer
type ArbitraryInt struct {
	// Digits stored from least significant to most significant
	// This allows easier manipulation of large numbers
	digits []int
	// Sign of the number (true for negative, false for positive)
	negative bool
}

// NewArbitraryInt creates a new ArbitraryInt from a string input
func NewArbitraryInt(input string) (*ArbitraryInt, error) {
	// Remove any whitespace
	input = strings.TrimSpace(input)

	// Check for sign
	negative := false
	if input[0] == '-' {
		negative = true
		input = input[1:]
	} else if input[0] == '+' {
		input = input[1:]
	}

	// Validate input contains only digits
	for _, char := range input {
		if char < '0' || char > '9' {
			return nil, fmt.Errorf("invalid input: %s", input)
		}
	}

	// Create digits slice (least significant first)
	digits := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		digits[len(input)-1-i] = int(input[i] - '0')
	}

	return &ArbitraryInt{
		digits:   digits,
		negative: negative,
	}, nil
}
