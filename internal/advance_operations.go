package internal

import "fmt"

// Pow computes the power of the number to a given exponent
func (a *ArbitraryInt) Pow(exponent *ArbitraryInt) (*ArbitraryInt, error) {
	// Handle special cases
	if exponent.negative {
		return nil, fmt.Errorf("negative exponents not supported")
	}

	// Convert exponent to int for easier computation
	expInt := 0
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
