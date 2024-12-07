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

// String converts ArbitraryInt back to a string representation
func (a *ArbitraryInt) String() string {
	// Handle zero case
	if len(a.digits) == 0 {
		return "0"
	}

	// Build string from most significant digit
	builder := strings.Builder{}

	// Add negative sign if needed
	if a.negative {
		builder.WriteRune('-')
	}

	// Write digits from most to least significant
	for i := len(a.digits) - 1; i >= 0; i-- {
		builder.WriteRune(rune(a.digits[i] + '0'))
	}

	return builder.String()
}

// Copy creates a deep copy of the ArbitraryInt
func (a *ArbitraryInt) Copy() *ArbitraryInt {
	digits := make([]int, len(a.digits))
	copy(digits, a.digits)
	return &ArbitraryInt{
		digits:   digits,
		negative: a.negative,
	}
}

// removeLeadingZeros removes unnecessary leading zeros
func (a *ArbitraryInt) removeLeadingZeros() {
	for len(a.digits) > 0 && a.digits[len(a.digits)-1] == 0 {
		a.digits = a.digits[:len(a.digits)-1]
	}

	// Handle zero case
	if len(a.digits) == 0 {
		a.negative = false
	}
}

// Abs returns the absolute value of the number
func (a *ArbitraryInt) Abs() *ArbitraryInt {
	copy := a.Copy()
	copy.negative = false
	return copy
}

// IsZero checks if the number is zero
func (a *ArbitraryInt) IsZero() bool {
	return len(a.digits) == 0
}
