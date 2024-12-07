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

// Add performs addition of two ArbitraryInt numbers
func (a *ArbitraryInt) Add(b *ArbitraryInt) *ArbitraryInt {
	// Handle sign differences
	if a.negative != b.negative {
		// If signs are different, this becomes a subtraction problem
		if a.negative {
			// -a + b = b - |a|
			return b.Subtract(a.Abs())
		}
		// a + (-b) = a - |b|
		return a.Subtract(b.Abs())
	}

	// Determine max length for addition
	maxLen := max(len(a.digits), len(b.digits))
	result := make([]int, maxLen+1) // +1 for potential carry

	// Perform digit-by-digit addition
	carry := 0
	for i := 0; i < maxLen; i++ {
		// Get digits, use 0 if index out of range
		aDigit := 0
		if i < len(a.digits) {
			aDigit = a.digits[i]
		}

		bDigit := 0
		if i < len(b.digits) {
			bDigit = b.digits[i]
		}

		// Add digits and carry
		sum := aDigit + bDigit + carry
		result[i] = sum % 10
		carry = sum / 10
	}

	// Handle final carry
	if carry > 0 {
		result[maxLen] = carry
	} else {
		result = result[:maxLen]
	}

	// Create new ArbitraryInt with the result
	return &ArbitraryInt{
		digits:   result,
		negative: a.negative, // Preserve original sign
	}
}

// Subtract performs subtraction of two ArbitraryInt numbers
func (a *ArbitraryInt) Subtract(b *ArbitraryInt) *ArbitraryInt {
	// Handle sign differences
	if a.negative != b.negative {
		// Different signs: a - (-b) = a + b
		// -a - b = -(a + b)
		result := a.Abs().Add(b.Abs())
		result.negative = a.negative
		return result
	}

	// Determine which absolute value is larger
	comparison := a.Abs().Compare(b.Abs())
	var larger, smaller *ArbitraryInt
	resultNegative := false

	switch comparison {
	case 0:
		// Equal magnitude, result is zero
		return &ArbitraryInt{digits: []int{}, negative: false}
	case 1:
		larger = a
		smaller = b
		resultNegative = a.negative
	case -1:
		larger = b
		smaller = a
		resultNegative = !a.negative
	}

	// Perform subtraction
	result := make([]int, len(larger.digits))
	copy(result, larger.digits)

	borrow := 0
	for i := 0; i < len(result); i++ {
		// Get smaller digit, use 0 if out of range
		smallerDigit := 0
		if i < len(smaller.digits) {
			smallerDigit = smaller.digits[i]
		}

		// Subtract with borrow
		currentDigit := result[i] - smallerDigit - borrow
		if currentDigit < 0 {
			currentDigit += 10
			borrow = 1
		} else {
			borrow = 0
		}

		result[i] = currentDigit
	}

	// Remove leading zeros
	for len(result) > 0 && result[len(result)-1] == 0 {
		result = result[:len(result)-1]
	}

	return &ArbitraryInt{
		digits:   result,
		negative: resultNegative,
	}
}

// Multiply performs multiplication of two ArbitraryInt numbers
func (a *ArbitraryInt) Multiply(b *ArbitraryInt) *ArbitraryInt {
	// Handle zero cases
	if a.IsZero() || b.IsZero() {
		return &ArbitraryInt{digits: []int{}, negative: false}
	}

	// Determine result sign
	resultNegative := a.negative != b.negative

	// Initialize result array
	result := make([]int, len(a.digits)+len(b.digits))

	// Perform long multiplication
	for i := 0; i < len(a.digits); i++ {
		for j := 0; j < len(b.digits); j++ {
			// Multiply digits and add to appropriate position
			product := a.digits[i] * b.digits[j]
			total := result[i+j] + product

			// Update result with current digit
			result[i+j] = total % 10

			// Carry over to next digit
			result[i+j+1] += total / 10
		}
	}

	// Remove leading zeros
	for len(result) > 0 && result[len(result)-1] == 0 {
		result = result[:len(result)-1]
	}

	return &ArbitraryInt{
		digits:   result,
		negative: resultNegative,
	}
}

// Divide performs division of two ArbitraryInt numbers
func (a *ArbitraryInt) Divide(b *ArbitraryInt) (*ArbitraryInt, *ArbitraryInt, error) {
	// Handle division by zero
	if b.IsZero() {
		return nil, nil, fmt.Errorf("division by zero")
	}

	// Handle zero dividend
	if a.IsZero() {
		return &ArbitraryInt{digits: []int{}, negative: false},
			&ArbitraryInt{digits: []int{}, negative: false},
			nil
	}

	// Determine result sign
	resultNegative := a.negative != b.negative

	// Take absolute values for division
	dividend := a.Abs()
	divisor := b.Abs()

	// If divisor is larger than dividend, quotient is zero, remainder is dividend
	if dividend.Compare(divisor) < 0 {
		return &ArbitraryInt{digits: []int{}, negative: false},
			a.Copy(),
			nil
	}

	// Long division algorithm
	quotient := &ArbitraryInt{digits: []int{}}
	remainder := &ArbitraryInt{digits: []int{}}

	// Start from the most significant digit
	for i := len(dividend.digits) - 1; i >= 0; i-- {
		// Build remainder
		remainder.digits = append([]int{dividend.digits[i]}, remainder.digits...)
		remainder.removeLeadingZeros()

		// Initialize current quotient digit
		currentQuotientDigit := 0

		// Repeatedly subtract divisor
		for remainder.Compare(divisor) >= 0 {
			remainder = remainder.Subtract(divisor)
			currentQuotientDigit++
		}

		// Prepend quotient digit
		if currentQuotientDigit > 0 || len(quotient.digits) > 0 {
			quotient.digits = append([]int{currentQuotientDigit}, quotient.digits...)
		}
	}

	// Set signs
	quotient.negative = resultNegative
	remainder.negative = a.negative // Remainder takes sign of dividend

	return quotient, remainder, nil
}

// Modulo returns the remainder of division
func (a *ArbitraryInt) Modulo(b *ArbitraryInt) (*ArbitraryInt, error) {
	_, remainder, err := a.Divide(b)
	return remainder, err
}

// Helper function to get maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
