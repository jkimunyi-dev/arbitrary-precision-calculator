package internal

import (
	"fmt"
	"math/big"
	"testing"
)

// TestNewArbitraryInt tests the creation of ArbitraryInt
func TestNewArbitraryInt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		error    bool
	}{
		{"123", "123", false},
		{"   123", "123", false},
		{"-000123", "-123", false},
		{"0", "0", false},
		{"", "", true},
		{"abc", "", true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Input: %v", tt.input), func(t *testing.T) {
			ai, err := NewArbitraryInt(tt.input)
			if (err != nil) != tt.error {
				t.Fatalf("unexpected error status: %v", err)
			}
			if !tt.error && ai.String() != tt.expected {
				t.Fatalf("expected %s, got %s", tt.expected, ai.String())
			}
		})
	}
}

// TestAddition tests addition of ArbitraryInt
func TestAddition(t *testing.T) {
	testCases := []struct {
		a, b     string
		expected string
	}{
		{"123", "456", "579"},
		{"-123", "456", "333"},
		{"123", "-456", "-333"},
		{"-123", "-456", "-579"},
		{"0", "0", "0"},
		{"999999", "1", "1000000"},
		{"1000000", "-999999", "1"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s + %s", tc.a, tc.b), func(t *testing.T) {
			numA, _ := NewArbitraryInt(tc.a)
			numB, _ := NewArbitraryInt(tc.b)

			result := numA.Add(numB)
			if result.String() != tc.expected {
				t.Errorf("Addition failed. Expected %s, got %s", tc.expected, result.String())
			}
		})
	}
}

// TestSubtraction tests subtraction of ArbitraryInt
func TestSubtraction(t *testing.T) {
	testCases := []struct {
		a, b     string
		expected string
	}{
		{"456", "123", "333"},
		{"123", "456", "-333"},
		{"-123", "-456", "333"},
		{"456", "-123", "579"},
		{"0", "0", "0"},
		{"1000000", "1", "999999"},
		{"1", "1000000", "-999999"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s - %s", tc.a, tc.b), func(t *testing.T) {
			numA, _ := NewArbitraryInt(tc.a)
			numB, _ := NewArbitraryInt(tc.b)

			result := numA.Subtract(numB)
			if result.String() != tc.expected {
				t.Errorf("Subtraction failed. Expected %s, got %s", tc.expected, result.String())
			}
		})
	}
}

// TestMultiplication tests multiplication of ArbitraryInt
func TestMultiplication(t *testing.T) {
	testCases := []struct {
		a, b     string
		expected string
	}{
		{"123", "456", "56088"},
		{"-123", "456", "-56088"},
		{"123", "-456", "-56088"},
		{"-123", "-456", "56088"},
		{"0", "1000000", "0"},
		{"999999", "999999", "999998000001"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s * %s", tc.a, tc.b), func(t *testing.T) {
			numA, _ := NewArbitraryInt(tc.a)
			numB, _ := NewArbitraryInt(tc.b)

			result := numA.Multiply(numB)
			if result.String() != tc.expected {
				t.Errorf("Multiplication failed. Expected %s, got %s", tc.expected, result.String())
			}
		})
	}
}

// TestDivision tests division of ArbitraryInt
func TestDivision(t *testing.T) {
	testCases := []struct {
		a, b           string
		expectedQuot   string
		expectedRem    string
		expectDivError bool
	}{
		{"456", "123", "3", "87", false},
		{"1000", "10", "100", "0", false},
		{"7", "3", "2", "1", false},
		{"123", "0", "", "", true},
		{"-456", "123", "-3", "-87", false},
		{"456", "-123", "-3", "87", false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s / %s", tc.a, tc.b), func(t *testing.T) {
			numA, _ := NewArbitraryInt(tc.a)
			numB, _ := NewArbitraryInt(tc.b)

			quotient, remainder, err := numA.Divide(numB)

			if tc.expectDivError {
				if err == nil {
					t.Errorf("Expected division error, got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected division error: %v", err)
				return
			}

			if quotient.String() != tc.expectedQuot {
				t.Errorf("Division quotient failed. Expected %s, got %s", tc.expectedQuot, quotient.String())
			}

			if remainder.String() != tc.expectedRem {
				t.Errorf("Division remainder failed. Expected %s, got %s", tc.expectedRem, remainder.String())
			}
		})
	}
}

// TestPower tests the Pow operation of ArbitraryInt
func TestPower(t *testing.T) {
	testCases := []struct {
		base     string
		exp      string
		expected string
	}{
		{"2", "10", "1024"},
		{"3", "4", "81"},
		{"10", "3", "1000"},
		{"5", "0", "1"},
		{"-2", "3", "-8"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s^%s", tc.base, tc.exp), func(t *testing.T) {
			base, _ := NewArbitraryInt(tc.base)
			exp, _ := NewArbitraryInt(tc.exp)

			result, err := base.Pow(exp)
			if err != nil {
				t.Errorf("Unexpected error in power calculation: %v", err)
				return
			}

			if result.String() != tc.expected {
				t.Errorf("Power calculation failed. Expected %s, got %s", tc.expected, result.String())
			}
		})
	}
}

// TestFactorial tests the Factorial operation of ArbitraryInt
func TestFactorial(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"0", "1"},
		{"1", "1"},
		{"5", "120"},
		{"10", "3628800"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s!", tc.input), func(t *testing.T) {
			num, _ := NewArbitraryInt(tc.input)

			result, err := num.Factorial()
			if err != nil {
				t.Errorf("Unexpected error in factorial calculation: %v", err)
				return
			}

			if result.String() != tc.expected {
				t.Errorf("Factorial calculation failed. Expected %s, got %s", tc.expected, result.String())
			}
		})
	}
}

// TestBaseConversion tests the BaseConvert operation of ArbitraryInt
func TestBaseConversion(t *testing.T) {
	// Prepare the test cases by handling errors from NewArbitraryInt
	tests := []struct {
		input       *ArbitraryInt
		fromBase    int
		toBase      int
		expectedStr string
		error       bool
	}{
		{mustNewArbitraryInt("10"), 10, 2, "1010", false},
		{mustNewArbitraryInt("1010"), 2, 10, "10", false},
		{mustNewArbitraryInt("0xFF"), 16, 10, "597", false},
		{nil, 10, 2, "", true},
		{mustNewArbitraryInt("123"), 1, 10, "", true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Convert %v from base %v to base %v", tt.input, tt.fromBase, tt.toBase), func(t *testing.T) {
			if tt.input == nil && !tt.error {
				t.Fatal("invalid test case: nil input but error expected to be false")
			}

			if tt.input != nil {
				result, err := tt.input.BaseConvert(tt.fromBase, tt.toBase)
				if (err != nil) != tt.error {
					t.Fatalf("unexpected error status: %v", err)
				}
				if !tt.error && result.String() != tt.expectedStr {
					t.Fatalf("expected %s, got %s", tt.expectedStr, result.String())
				}
			}
		})
	}
}

// Helper function to create an ArbitraryInt or fail the test immediately.
func mustNewArbitraryInt(value string) *ArbitraryInt {
	ai, err := NewArbitraryInt(value)
	if err != nil {
		panic(fmt.Sprintf("failed to create ArbitraryInt: %v", err))
	}
	return ai
}

// TestLargeNumberOperations tests operations with very large numbers
func TestLargeNumberOperations(t *testing.T) {
	largeNumberCases := []struct {
		a, b      string
		operation string
		expected  string
	}{
		{
			"123456",
			"987654",
			"add",
			"1111110",
		},
		{
			"987654",
			"123456",
			"subtract",
			"864198",
		},
		{
			"123456",
			"987654",
			"multiply",
			"121931812224",
		},
	}

	for _, tc := range largeNumberCases {
		t.Run(fmt.Sprintf("Large number %s", tc.operation), func(t *testing.T) {
			numA, _ := NewArbitraryInt(tc.a)
			numB, _ := NewArbitraryInt(tc.b)

			var result *ArbitraryInt
			// var err error

			switch tc.operation {
			case "add":
				result = numA.Add(numB)
			case "subtract":
				result = numA.Subtract(numB)
			case "multiply":
				result = numA.Multiply(numB)
			default:
				t.Fatalf("Unsupported operation: %s", tc.operation)
			}

			if result.String() != tc.expected {
				t.Errorf("%s failed. Expected %s, got %s",
					tc.operation, tc.expected, result.String())
			}
		})
	}
}

// BenchmarkAddition benchmarks the addition operation
func BenchmarkAddition(b *testing.B) {
	num1, errA := NewArbitraryInt("123456789012345678901234567890")
	num2, errB := NewArbitraryInt("987654321098765432109876543210")

	if errA != nil || errB != nil {
		b.Fatalf("Failed to create test numbers: %v, %v", errA, errB)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create copies to avoid mutating original numbers
		copyA := num1.Copy()
		copyB := num2.Copy()
		copyA.Add(copyB)
	}
}

// BenchmarkMultiplication benchmarks the multiplication operation
func BenchmarkMultiplication(b *testing.B) {
	num1, errA := NewArbitraryInt("123456789012345678901234567890")
	num2, errB := NewArbitraryInt("987654321098765432109876543210")

	if errA != nil || errB != nil {
		b.Fatalf("Failed to create test numbers: %v, %v", errA, errB)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create copies to avoid mutating original numbers
		copyA := num1.Copy()
		copyB := num2.Copy()
		copyA.Multiply(copyB)
	}
}

// CompareWithBigInt provides a validation method against Go's big.Int
func TestCompareWithBigInt(t *testing.T) {
	testCases := []struct {
		a, b      string
		operation string
	}{
		{"123456789", "987654321", "add"},
		{"987654321", "123456789", "subtract"},
		{"54321", "12345", "multiply"},
		{"1000000", "3", "divide"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s %s %s", tc.a, tc.operation, tc.b), func(t *testing.T) {
			bigA, _ := new(big.Int).SetString(tc.a, 10)
			bigB, _ := new(big.Int).SetString(tc.b, 10)

			arbA, errA := NewArbitraryInt(tc.a)
			arbB, errB := NewArbitraryInt(tc.b)

			if errA != nil || errB != nil {
				t.Fatalf("Failed to create test numbers: %v, %v", errA, errB)
			}

			var arbResult *ArbitraryInt
			var bigResult *big.Int

			switch tc.operation {
			case "add":
				arbResult = arbA.Add(arbB)
				bigResult = new(big.Int).Add(bigA, bigB)
			case "subtract":
				arbResult = arbA.Subtract(arbB)
				bigResult = new(big.Int).Sub(bigA, bigB)
			case "multiply":
				arbResult = arbA.Multiply(arbB)
				bigResult = new(big.Int).Mul(bigA, bigB)
			case "divide":
				quotient, _, err := arbA.Divide(arbB)
				if err != nil {
					t.Fatalf("Division error: %v", err)
				}
				arbResult = quotient
				bigResult = new(big.Int).Div(bigA, bigB)
			default:
				t.Fatalf("Unsupported operation: %s", tc.operation)
			}

			if arbResult.String() != bigResult.String() {
				t.Errorf("%s operation failed. ArbitraryInt: %s, big.Int: %s",
					tc.operation, arbResult.String(), bigResult.String())
			}
		})
	}
}
