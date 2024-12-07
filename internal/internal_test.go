package internal

import (
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	numA, _ := NewArbitraryInt("123456789012345678901234567890")
	numB, _ := NewArbitraryInt("987654321098765432109876543210")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		numA.Add(numB)
	}
}

func BenchmarkMultiply(b *testing.B) {
	numA, _ := NewArbitraryInt("123456789")
	numB, _ := NewArbitraryInt("987654321")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		numA.Multiply(numB)
	}
}

func BenchmarkDivide(b *testing.B) {
	numA, _ := NewArbitraryInt("123456789012345678901234567890")
	numB, _ := NewArbitraryInt("987654321")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		numA.Divide(numB)
	}
}

func TestArithmeticOperations(t *testing.T) {
	// Reduced test cases with focus on key scenarios
	additionTests := []struct {
		a, b     string
		expected string
	}{
		{"123", "456", "579"},
		{"999", "1", "1000"},
		{"-123", "456", "333"},
	}

	subtractionTests := []struct {
		a, b     string
		expected string
	}{
		{"456", "123", "333"},
		{"1000", "1", "999"},
		{"-123", "456", "-579"},
	}

	multiplicationTests := []struct {
		a, b     string
		expected string
	}{
		{"123", "456", "56088"},
		{"999", "999", "998001"},
		{"-123", "456", "-56088"},
	}

	divisionTests := []struct {
		a, b         string
		expectedQuot string
		expectedRem  string
	}{
		{"10", "2", "5", "0"},
		{"15", "4", "3", "3"},
		{"100", "10", "10", "0"},
	}

	// Run addition tests
	t.Run("Addition", func(t *testing.T) {
		for _, tc := range additionTests {
			numA, _ := NewArbitraryInt(tc.a)
			numB, _ := NewArbitraryInt(tc.b)

			result := numA.Add(numB)
			if result.String() != tc.expected {
				t.Errorf("%s + %s: Expected %s, got %s", tc.a, tc.b, tc.expected, result.String())
			}
		}
	})

	// Run subtraction tests
	t.Run("Subtraction", func(t *testing.T) {
		for _, tc := range subtractionTests {
			numA, _ := NewArbitraryInt(tc.a)
			numB, _ := NewArbitraryInt(tc.b)

			result := numA.Subtract(numB)
			if result.String() != tc.expected {
				t.Errorf("%s - %s: Expected %s, got %s", tc.a, tc.b, tc.expected, result.String())
			}
		}
	})

	// Run multiplication tests
	t.Run("Multiplication", func(t *testing.T) {
		for _, tc := range multiplicationTests {
			numA, _ := NewArbitraryInt(tc.a)
			numB, _ := NewArbitraryInt(tc.b)

			result := numA.Multiply(numB)
			if result.String() != tc.expected {
				t.Errorf("%s * %s: Expected %s, got %s", tc.a, tc.b, tc.expected, result.String())
			}
		}
	})

	// Run division tests
	t.Run("Division", func(t *testing.T) {
		for _, tc := range divisionTests {
			numA, _ := NewArbitraryInt(tc.a)
			numB, _ := NewArbitraryInt(tc.b)

			quotient, remainder, err := numA.Divide(numB)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				continue
			}

			if quotient.String() != tc.expectedQuot {
				t.Errorf("%s ÷ %s: Expected quotient %s, got %s",
					tc.a, tc.b, tc.expectedQuot, quotient.String())
			}

			if remainder.String() != tc.expectedRem {
				t.Errorf("%s ÷ %s: Expected remainder %s, got %s",
					tc.a, tc.b, tc.expectedRem, remainder.String())
			}
		}
	})
}
