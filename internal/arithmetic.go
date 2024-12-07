package internal

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

// ArbitraryInt represents an arbitrary precision integer
type ArbitraryInt struct {
	// Digits stored from least significant to most significant
	digits []uint32
	// Sign of the number (true for negative, false for positive)
	negative bool
}

// NewArbitraryInt creates a new ArbitraryInt from a string input
func NewArbitraryInt(s string) (*ArbitraryInt, error) {
	// Remove any whitespace
	s = strings.TrimSpace(s)

	// Check for empty string
	if s == "" {
		return nil, fmt.Errorf("cannot create ArbitraryInt from empty string")
	}

	// Determine base based on prefix
	var base int
	if len(s) > 2 && s[:2] == "0x" { // Hexadecimal
		base = 16
		s = s[2:] // Remove the "0x" prefix
	} else if len(s) > 2 && s[:2] == "0b" { // Binary
		base = 2
		s = s[2:] // Remove the "0b" prefix
	} else { // Default to decimal
		base = 10
	}

	// Check for valid format
	valid := false
	switch base {
	case 10:
		valid, _ = regexp.MatchString(`^-?\d+$`, s)
	case 16:
		valid, _ = regexp.MatchString(`^-?[0-9a-fA-F]+$`, s)
	case 2:
		valid, _ = regexp.MatchString(`^-?[01]+$`, s)
	}
	if !valid {
		return nil, fmt.Errorf("invalid number format for base %d: %s", base, s)
	}

	// Determine sign
	negative := s[0] == '-'
	if negative {
		s = s[1:]
	}

	// Remove leading zeros
	s = strings.TrimLeft(s, "0")
	if s == "" {
		s = "0"
	}

	// Convert to base-10 representation
	var decimalStr string
	if base != 10 {
		val := new(big.Int)
		_, success := val.SetString(s, base) // Parse based on base
		if !success {
			return nil, fmt.Errorf("failed to parse number: %s", s)
		}
		decimalStr = val.Text(10) // Convert to decimal string
	} else {
		decimalStr = s
	}

	// Optimize digit storage using uint32 to pack more digits
	const digitBase = 1_000_000_000 // 9 digits per uint32
	digits := make([]uint32, (len(decimalStr)+8)/9)

	for i := 0; i < len(decimalStr); i += 9 {
		end := i + 9
		if end > len(decimalStr) {
			end = len(decimalStr)
		}
		chunk := decimalStr[i:end]

		// Convert chunk to uint32
		value := uint32(0)
		for _, digit := range chunk {
			value = value*10 + uint32(digit-'0')
		}

		digits[len(digits)-1-i/9] = value
	}

	return &ArbitraryInt{
		digits:   digits,
		negative: negative,
	}, nil
}

// Multiply implements efficient multiplication using Karatsuba algorithm
func (a *ArbitraryInt) Multiply(b *ArbitraryInt) *ArbitraryInt {
	if a == nil || b == nil {
		panic("cannot multiply nil ArbitraryInt")
	}

	// Handle zero cases
	if len(a.digits) == 0 || len(b.digits) == 0 {
		return &ArbitraryInt{digits: []uint32{0}}
	}

	// Determine sign
	negative := a.negative != b.negative

	// Use Karatsuba multiplication for large numbers
	if len(a.digits) > 32 || len(b.digits) > 32 {
		return a.karatsubaMultiply(b)
	}

	// Fallback to standard long multiplication for smaller numbers
	return a.longMultiply(b, negative)
}

// longMultiply performs standard long multiplication
func (a *ArbitraryInt) longMultiply(b *ArbitraryInt, negative bool) *ArbitraryInt {
	const base = uint64(1_000_000_000)
	result := make([]uint32, len(a.digits)+len(b.digits))

	for i, multiplier := range a.digits {
		if multiplier == 0 {
			continue
		}

		carry := uint64(0)
		for j, multiplicand := range b.digits {
			product := uint64(multiplier)*uint64(multiplicand) + uint64(result[i+j]) + carry
			result[i+j] = uint32(product % base)
			carry = product / base
		}

		if carry > 0 {
			result[i+len(b.digits)] += uint32(carry)
		}
	}

	// Trim leading zeros
	for len(result) > 0 && result[len(result)-1] == 0 {
		result = result[:len(result)-1]
	}

	if len(result) == 0 {
		result = []uint32{0}
	}

	return &ArbitraryInt{
		digits:   result,
		negative: negative,
	}
}

// karatsubaMultiply implements Karatsuba multiplication algorithm
func (a *ArbitraryInt) karatsubaMultiply(b *ArbitraryInt) *ArbitraryInt {
	// Determine sign
	negative := a.negative != b.negative

	// If numbers are small, use long multiplication
	if len(a.digits) <= 32 || len(b.digits) <= 32 {
		return a.longMultiply(b, negative)
	}

	// Determine split point
	m := min(len(a.digits), len(b.digits)) / 2

	// Split numbers
	aLow := &ArbitraryInt{digits: a.digits[:m], negative: false}
	aHigh := &ArbitraryInt{digits: a.digits[m:], negative: false}
	bLow := &ArbitraryInt{digits: b.digits[:m], negative: false}
	bHigh := &ArbitraryInt{digits: b.digits[m:], negative: false}

	// Recursive Karatsuba steps
	z0 := aLow.Multiply(bLow)
	z2 := aHigh.Multiply(bHigh)

	// Intermediate calculation
	lowSum := aLow.Add(aHigh)
	highSum := bLow.Add(bHigh)
	z1 := lowSum.Multiply(highSum).Subtract(z0).Subtract(z2)

	// Combine results with proper scaling
	result := z0
	z1Scaled := z1
	z2Scaled := z2

	// Complex result combination
	for i := 0; i < 2*m; i++ {
		result = result.leftShift()
	}
	for i := 0; i < m; i++ {
		z1Scaled = z1Scaled.leftShift()
	}

	result = result.Add(z1Scaled).Add(z2Scaled)
	result.negative = negative

	return result
}

// leftShift multiplies the number by base (essentially adding zeros)
func (a *ArbitraryInt) leftShift() *ArbitraryInt {
	if len(a.digits) == 1 && a.digits[0] == 0 {
		return a
	}

	result := make([]uint32, len(a.digits)+1)
	copy(result[1:], a.digits)
	return &ArbitraryInt{
		digits:   result,
		negative: a.negative,
	}
}

// Add performs efficient addition
func (a *ArbitraryInt) Add(b *ArbitraryInt) *ArbitraryInt {
	if a.negative && !b.negative {
		a.negative = false
		result := b.Subtract(a)
		return result
	}
	if !a.negative && b.negative {
		b.negative = false
		return a.Subtract(b)
	}

	maxLen := max(len(a.digits), len(b.digits))
	result := make([]uint32, maxLen+1) // Extra space for carry
	carry := uint32(0)

	const digitBase = 1_000_000_000 // Base used for digit packing, 9 digits per uint32

	// Zero-padding a.digits and b.digits
	aDigits := make([]uint32, maxLen)
	bDigits := make([]uint32, maxLen)
	copy(aDigits[:len(a.digits)], a.digits)
	copy(bDigits[:len(b.digits)], b.digits)

	// Perform addition with carry
	for i := 0; i < maxLen; i++ {
		sum := aDigits[i] + bDigits[i] + carry
		result[i] = sum % digitBase
		carry = sum / digitBase
	}

	// Handle final carry
	if carry > 0 {
		result[maxLen] = carry
	} else {
		result = result[:maxLen] // Trim unused space if no carry
	}

	// Trim leading zeros
	for len(result) > 1 && result[len(result)-1] == 0 {
		result = result[:len(result)-1]
	}

	return &ArbitraryInt{
		digits:   result,
		negative: a.negative,
	}
}

// Subtract performs efficient subtraction
func (a *ArbitraryInt) Subtract(b *ArbitraryInt) *ArbitraryInt {
	// Handle sign differences
	if a.negative != b.negative {
		b.negative = !b.negative
		result := a.Add(b)
		b.negative = !b.negative
		return result
	}

	// Determine if we need to swap and track original sign
	swap := false
	if len(a.digits) < len(b.digits) ||
		(len(a.digits) == len(b.digits) && compareDigits(a.digits, b.digits) < 0) {
		a, b = b, a
		swap = true
	}

	result := make([]uint32, len(a.digits))
	copy(result, a.digits)

	borrow := uint32(0)
	for i := 0; i < len(b.digits); i++ {
		if result[i] < b.digits[i]+borrow {
			result[i] = result[i] + 1_000_000_000 - b.digits[i] - borrow
			borrow = 1
		} else {
			result[i] -= b.digits[i] + borrow
			borrow = 0
		}
	}

	// Propagate borrow
	for i := len(b.digits); borrow > 0 && i < len(result); i++ {
		if result[i] == 0 {
			result[i] = 999_999_999
		} else {
			result[i]--
			borrow = 0
		}
	}

	// Trim leading zeros
	for len(result) > 1 && result[len(result)-1] == 0 {
		result = result[:len(result)-1]
	}

	// Determine final sign
	negative := a.negative
	if swap {
		negative = !negative
	}

	return &ArbitraryInt{
		digits:   result,
		negative: negative,
	}
}

// compareDigits compares two digit arrays
func compareDigits(a, b []uint32) int {
	if len(a) != len(b) {
		return len(a) - len(b)
	}

	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != b[i] {
			return int(a[i]) - int(b[i])
		}
	}
	return 0
}

// Copy creates a deep copy of the ArbitraryInt
func (a *ArbitraryInt) Copy() *ArbitraryInt {
	digits := make([]uint32, len(a.digits))
	copy(digits, a.digits)

	return &ArbitraryInt{
		digits:   digits,
		negative: a.negative,
	}
}

// String converts ArbitraryInt to string representation
func (a *ArbitraryInt) String() string {
	if len(a.digits) == 0 || (len(a.digits) == 1 && a.digits[0] == 0) {
		return "0"
	}

	// Build string representation
	var parts []string
	for i := len(a.digits) - 1; i >= 0; i-- {
		if i == len(a.digits)-1 {
			parts = append(parts, fmt.Sprintf("%d", a.digits[i]))
		} else {
			parts = append(parts, fmt.Sprintf("%09d", a.digits[i]))
		}
	}
	result := strings.Join(parts, "")
	if a.negative {
		result = "-" + result
	}
	return result
}

// Utility function to find min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Utility function to find max
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (a *ArbitraryInt) Abs() *ArbitraryInt {
	abs := &ArbitraryInt{
		digits:   make([]uint32, len(a.digits)),
		negative: false,
	}
	copy(abs.digits, a.digits)
	return abs
}

func compareAbsolute(a, b *ArbitraryInt) int {
	// Compare lengths first
	if len(a.digits) > len(b.digits) {
		return 1
	}
	if len(a.digits) < len(b.digits) {
		return -1
	}

	// Compare digit by digit from most significant
	for i := len(a.digits) - 1; i >= 0; i-- {
		if a.digits[i] > b.digits[i] {
			return 1
		}
		if a.digits[i] < b.digits[i] {
			return -1
		}
	}
	return 0
}

func subtractAbsolute(a, b *ArbitraryInt) *ArbitraryInt {
	result := &ArbitraryInt{digits: []uint32{}}
	borrow := uint32(0)

	// Ensure a is larger
	if compareAbsolute(a, b) < 0 {
		a, b = b, a
	}

	for i := 0; i < len(a.digits); i++ {
		var current uint32
		if i < len(b.digits) {
			current = a.digits[i] - b.digits[i] - borrow
		} else {
			current = a.digits[i] - borrow
		}

		if current > a.digits[i] {
			current += 1_000_000_000
			borrow = 1
		} else {
			borrow = 0
		}

		result.digits = append(result.digits, current)
	}

	return trimLeadingZeros(result)
}

func trimLeadingZeros(a *ArbitraryInt) *ArbitraryInt {
	if len(a.digits) == 0 {
		return &ArbitraryInt{digits: []uint32{0}}
	}

	for len(a.digits) > 1 && a.digits[len(a.digits)-1] == 0 {
		a.digits = a.digits[:len(a.digits)-1]
	}

	return a
}

// Divide performs division of two ArbitraryInt numbers
func (a *ArbitraryInt) Divide(other *ArbitraryInt) (*ArbitraryInt, *ArbitraryInt, error) {
	// Check for division by zero
	if isZero(other) {
		return nil, nil, fmt.Errorf("division by zero")
	}

	// Handle sign for division
	negative := a.negative != other.negative

	// Convert to absolute values for division
	dividend := a.Abs()
	divisor := other.Abs()

	// Check if divisor is larger than dividend
	if compareAbsolute(dividend, divisor) < 0 {
		value, err := NewArbitraryInt("0")
		if err != nil {
			// Handle the error
			fmt.Println("Error creating ArbitraryInt:", err)
			return value, dividend, nil

		}
		// Use `value` here

	}

	// Long division algorithm
	quotient := &ArbitraryInt{digits: []uint32{}}
	remainder := &ArbitraryInt{digits: []uint32{}}

	for i := len(dividend.digits) - 1; i >= 0; i-- {
		// Build remainder
		remainder.digits = append([]uint32{dividend.digits[i]}, remainder.digits...)

		// Normalize remainder
		remainder = trimLeadingZeros(remainder)

		// Find how many times divisor goes into remainder
		multiple := uint32(0)
		for compareAbsolute(remainder, divisor) >= 0 {
			remainder = subtractAbsolute(remainder, divisor)
			multiple++
		}

		// Prepend multiple to quotient
		quotient.digits = append([]uint32{multiple}, quotient.digits...)
	}

	// Trim leading zeros
	quotient = trimLeadingZeros(quotient)
	remainder = trimLeadingZeros(remainder)

	// Set sign
	quotient.negative = negative
	remainder.negative = a.negative

	return quotient, remainder, nil
}

// Modulo calculates the remainder of division
func (a *ArbitraryInt) Modulo(other *ArbitraryInt) (*ArbitraryInt, error) {
	// Check for division by zero
	if isZero(other) {
		return nil, fmt.Errorf("modulo by zero")
	}

	// Implement modulo logic
	// This is a placeholder - you'll need to implement the full modulo algorithm
	result := &ArbitraryInt{}
	return result, nil
}

// Helper function to check if an ArbitraryInt is zero
func isZero(a *ArbitraryInt) bool {
	return len(a.digits) == 0 || (len(a.digits) == 1 && a.digits[0] == 0)
}
