package internal

import (
	"math/big"
	"testing"
)

// TestArbitraryIntBasicOperations tests basic arithmetic operations
func TestArbitraryIntBasicOperations(t *testing.T) {
	// Test cases for addition
	addTestCases := []struct {
		a, b     string
		expected string
	}{
		{"123", "456", "579"},
		{"0", "0", "0"},
		{"-100", "50", "-50"},
		{"999999", "1", "1000000"},
	}

	for _, tc := range addTestCases {
		t.Run(tc.a+" + "+tc.b, func(t *testing.T) {
			num1, err1 := NewArbitraryInt(tc.a)
			num2, err2 := NewArbitraryInt(tc.b)

			if err1 != nil || err2 != nil {
				t.Fatalf("Failed to create ArbitraryInt: err1=%v, err2=%v", err1, err2)
			}

			result := num1.Add(num2)
			if result.String() != tc.expected {
				t.Errorf("Add(%s, %s): got %v, want %v", tc.a, tc.b, result, tc.expected)
			}
		})
	}

	// Test cases for subtraction
	subtractTestCases := []struct {
		a, b     string
		expected string
	}{
		{"500", "300", "200"},
		{"0", "0", "0"},
		{"-100", "50", "-150"},
		{"1000", "1", "999"},
	}

	for _, tc := range subtractTestCases {
		t.Run(tc.a+" - "+tc.b, func(t *testing.T) {
			num1, err1 := NewArbitraryInt(tc.a)
			num2, err2 := NewArbitraryInt(tc.b)

			if err1 != nil || err2 != nil {
				t.Fatalf("Failed to create ArbitraryInt: err1=%v, err2=%v", err1, err2)
			}

			result := num1.Subtract(num2)
			if result.String() != tc.expected {
				t.Errorf("Subtract(%s, %s): got %v, want %v", tc.a, tc.b, result, tc.expected)
			}
		})
	}

	// Test cases for multiplication
	multiplyTestCases := []struct {
		a, b     string
		expected string
	}{
		{"10", "20", "200"},
		{"0", "1000", "0"},
		{"-5", "7", "-35"},
		{"999999", "999999", "999998000001"},
	}

	for _, tc := range multiplyTestCases {
		t.Run(tc.a+" * "+tc.b, func(t *testing.T) {
			num1, err1 := NewArbitraryInt(tc.a)
			num2, err2 := NewArbitraryInt(tc.b)

			if err1 != nil || err2 != nil {
				t.Fatalf("Failed to create ArbitraryInt: err1=%v, err2=%v", err1, err2)
			}

			result := num1.Multiply(num2)
			if result.String() != tc.expected {
				t.Errorf("Multiply(%s, %s): got %v, want %v", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}

// TestLargeNumberOperations tests operations with extremely large numbers
func TestLargeNumberOperations(t *testing.T) {
	// Large number test cases
	largeNumberTestCases := []struct {
		a, b   string
		addExp string
		subExp string
		mulExp string
	}{
		{
			"123456789012345678901234567890",
			"987654321098765432109876543210",
			"1111111110111111111011111111100",
			"-864197532086419753208641975320",
			"121932631124517947378194577901322723971934093692116447368900",
		},
	}

	for _, tc := range largeNumberTestCases {
		t.Run("Large Number Operations", func(t *testing.T) {
			num1, err1 := NewArbitraryInt(tc.a)
			num2, err2 := NewArbitraryInt(tc.b)

			if err1 != nil || err2 != nil {
				t.Fatalf("Failed to create ArbitraryInt: err1=%v, err2=%v", err1, err2)
			}

			// Addition test
			addResult := num1.Add(num2)
			if addResult.String() != tc.addExp {
				t.Errorf("Add large numbers: got %v, want %v", addResult, tc.addExp)
			}

			// Subtraction test
			subResult := num1.Subtract(num2)
			if subResult.String() != tc.subExp {
				t.Errorf("Subtract large numbers: got %v, want %v", subResult, tc.subExp)
			}

			// Multiplication test
			mulResult := num1.Multiply(num2)
			if mulResult.String() != tc.mulExp {
				t.Errorf("Multiply large numbers: got %v, want %v", mulResult, tc.mulExp)
			}
		})
	}
}

// TestDivisionAndModulo tests division and modulo operations
func TestDivisionAndModulo(t *testing.T) {
	divisionTestCases := []struct {
		a, b           string
		expectedQuot   string
		expectedRem    string
		expectingError bool
	}{
		{"10", "3", "3", "1", false},
		{"100", "10", "10", "0", false},
		{"7", "2", "3", "1", false},
		{"0", "5", "0", "0", false},
		// Division by zero test
		{"10", "0", "", "", true},
	}

	for _, tc := range divisionTestCases {
		t.Run(tc.a+" / "+tc.b, func(t *testing.T) {
			num1, err1 := NewArbitraryInt(tc.a)
			num2, err2 := NewArbitraryInt(tc.b)

			if err1 != nil || err2 != nil {
				t.Fatalf("Failed to create ArbitraryInt: err1=%v, err2=%v", err1, err2)
			}

			quotient, remainder, err := num1.Divide(num2)

			if tc.expectingError {
				if err == nil {
					t.Errorf("Expected error for division by zero")
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error in division: %v", err)
			}

			if quotient.String() != tc.expectedQuot {
				t.Errorf("Division quotient: got %v, want %v", quotient, tc.expectedQuot)
			}

			if remainder.String() != tc.expectedRem {
				t.Errorf("Division remainder: got %v, want %v", remainder, tc.expectedRem)
			}
		})
	}
}

// TestPowerOperation tests exponentiation
func TestPowerOperation(t *testing.T) {
	powerTestCases := []struct {
		base, exp string
		expected  string
	}{
		{"2", "10", "1024"},
		{"3", "4", "81"},
		{"10", "3", "1000"},
		{"0", "5", "0"},
		{"5", "0", "1"},
	}

	for _, tc := range powerTestCases {
		t.Run(tc.base+"^"+tc.exp, func(t *testing.T) {
			base, err1 := NewArbitraryInt(tc.base)
			exp, err2 := NewArbitraryInt(tc.exp)

			if err1 != nil || err2 != nil {
				t.Fatalf("Failed to create ArbitraryInt: err1=%v, err2=%v", err1, err2)
			}

			result, err := base.Pow(exp)
			if err != nil {
				t.Fatalf("Unexpected error in power operation: %v", err)
			}

			if result.String() != tc.expected {
				t.Errorf("Power(%s, %s): got %v, want %v", tc.base, tc.exp, result, tc.expected)
			}
		})
	}
}

// BenchmarkLargeNumberMultiplication benchmarks multiplication of large numbers
// BenchmarkLargeNumberMultiplication benchmarks multiplication of large numbers
// BenchmarkLargeNumberMultiplication benchmarks multiplication of large numbers
func BenchmarkLargeNumberMultiplication(b *testing.B) {
	// Prepare large numbers for benchmarking outside of the loop
	aStr := "123456789012345678901234567890"
	bStr := "987654321098765432109876543210"

	_, err1 := NewArbitraryInt(aStr)
	_, err2 := NewArbitraryInt(bStr)

	if err1 != nil || err2 != nil {
		b.Fatalf("Failed to create ArbitraryInt: err1=%v, err2=%v", err1, err2)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create new instances in each iteration to avoid mutation
		a, _ := NewArbitraryInt(aStr)
		b, _ := NewArbitraryInt(bStr)
		a.Multiply(b)
	}
}

// CompareWithStandardLibrary provides a comparison with math/big for verification
func TestCompareWithStandardLibrary(t *testing.T) {
	testCases := []struct {
		a, b      string
		operation string
	}{
		{"1234567890", "9876543210", "add"},
		{"9876543210", "1234567890", "subtract"},
		{"12345", "6789", "multiply"},
	}

	for _, tc := range testCases {
		t.Run(tc.a+" "+tc.operation+" "+tc.b, func(t *testing.T) {
			// Create ArbitraryInt
			num1, err1 := NewArbitraryInt(tc.a)
			num2, err2 := NewArbitraryInt(tc.b)

			if err1 != nil || err2 != nil {
				t.Fatalf("Failed to create ArbitraryInt: err1=%v, err2=%v", err1, err2)
			}

			// Create big.Int for comparison
			bigNum1 := new(big.Int)
			bigNum2 := new(big.Int)
			bigNum1.SetString(tc.a, 10)
			bigNum2.SetString(tc.b, 10)

			var arbitraryResult string
			var bigResult string

			switch tc.operation {
			case "add":
				arbitraryResult = num1.Add(num2).String()
				bigResult = bigNum1.Add(bigNum1, bigNum2).String()
			case "subtract":
				arbitraryResult = num1.Subtract(num2).String()
				bigResult = bigNum1.Sub(bigNum1, bigNum2).String()
			case "multiply":
				arbitraryResult = num1.Multiply(num2).String()
				bigResult = bigNum1.Mul(bigNum1, bigNum2).String()
			}

			if arbitraryResult != bigResult {
				t.Errorf("%s operation mismatch: ArbitraryInt=%s, big.Int=%s",
					tc.operation, arbitraryResult, bigResult)
			}
		})
	}
}
