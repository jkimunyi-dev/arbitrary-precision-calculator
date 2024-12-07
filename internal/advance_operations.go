package internal

import (
	"fmt"
	"math/big"
)

// Pow computes the power of the number to a given exponent
func (a *ArbitraryInt) Pow(exponent *ArbitraryInt) (*ArbitraryInt, error) {
	// Handle special cases
	if exponent.negative {
		return nil, fmt.Errorf("negative exponents not supported")
	}

	// Convert exponent to int for easier computation
	expInt := uint32(0)
	for _, digit := range exponent.digits {
		expInt = expInt*10 + digit
	}

	// Special cases for exponent
	if expInt == 0 {
		// Any number to the 0th power is 1
		return NewArbitraryInt("1")
	}
	if expInt == 1 {
		// Any number to the 1st power is itself
		return a.Copy(), nil
	}

	// Use exponentiation by squaring for efficiency
	result, _ := NewArbitraryInt("1")
	base := a.Copy()

	for expInt > 0 {
		// If the current bit is 1, multiply the result
		if expInt%2 == 1 {
			result = result.Multiply(base)
		}

		// Square the base
		base = base.Multiply(base)
		expInt /= 2
	}

	// Handle negative base with odd exponent
	result.negative = a.negative && (exponent.digits[0]%2 == 1)

	return result, nil
}

// Factorial computes the factorial of a non-negative integer
func (a *ArbitraryInt) Factorial() (*ArbitraryInt, error) {
	// Check for negative input
	if a.negative {
		return nil, fmt.Errorf("factorial is not defined for negative numbers")
	}

	// Convert to integer for easier computation
	factorialLimit := uint32(0)
	for _, digit := range a.digits {
		factorialLimit = factorialLimit*10 + digit
	}

	// Handle special cases
	if factorialLimit == 0 || factorialLimit == 1 {
		return NewArbitraryInt("1")
	}

	// Compute factorial using multiplication
	result, _ := NewArbitraryInt("1")
	current, _ := NewArbitraryInt("1")

	// Create a zero for comparison
	// zero, _ := NewArbitraryInt("0")

	// Use a wrapper for Add that extracts the first return value
	addOne, _ := NewArbitraryInt("1")

	for compareResult(current, a) <= 0 {
		result = result.Multiply(current)
		current = current.Add(addOne)
	}

	return result, nil
}

// BaseConvert converts the number from one base to another
func (a *ArbitraryInt) BaseConvert(fromBase, toBase int) (*ArbitraryInt, error) {
	if fromBase < 2 || toBase < 2 {
		return nil, fmt.Errorf("base must be >= 2")
	}

	// Convert current number to a string in base `fromBase`
	value := new(big.Int)
	base10Str := a.String() // Assumes `String()` provides the number in base 10
	_, success := value.SetString(base10Str, fromBase)
	if !success {
		return nil, fmt.Errorf("failed to convert from base %d", fromBase)
	}

	// Convert to target base
	targetStr := value.Text(toBase)

	// Create a new ArbitraryInt from the target base string
	return NewArbitraryInt(targetStr)
}

// Helper function to handle comparison with error return
func compareResult(a, b *ArbitraryInt) int {
	return a.Compare(b)
}

// Helper function to convert digit to character for base conversion
func convertDigitToChar(digit int) rune {
	if digit >= 0 && digit <= 9 {
		return rune('0' + digit)
	}
	return rune('A' + digit - 10)
}

// Compare compares two ArbitraryInt values
func (a *ArbitraryInt) Compare(b *ArbitraryInt) int {
	// First, compare signs
	if a.negative && !b.negative {
		return -1
	}
	if !a.negative && b.negative {
		return 1
	}

	// If signs are different, comparison is done
	// If both are negative, we'll invert the comparison
	signMultiplier := 1
	if a.negative {
		signMultiplier = -1
	}

	// Compare lengths
	if len(a.digits) > len(b.digits) {
		return 1 * signMultiplier
	}
	if len(a.digits) < len(b.digits) {
		return -1 * signMultiplier
	}

	// Compare digit by digit from most significant
	for i := len(a.digits) - 1; i >= 0; i-- {
		if a.digits[i] > b.digits[i] {
			return 1 * signMultiplier
		}
		if a.digits[i] < b.digits[i] {
			return -1 * signMultiplier
		}
	}

	// Numbers are equal
	return 0
}
