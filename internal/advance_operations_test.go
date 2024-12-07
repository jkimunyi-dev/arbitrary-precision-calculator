package internal

import (
	"testing"
)

// Benchmarks for different operation sizes
func BenchmarkAdd(b *testing.B) {
	benchmarkCases := []struct {
		name string
		a, b string
	}{
		{"SmallNumbers", "1234", "5678"},
		{"LargeNumbers", "123456789012345678901234567890", "987654321098765432109876543210"},
	}

	for _, bc := range benchmarkCases {
		b.Run(bc.name, func(b *testing.B) {
			numA, _ := NewArbitraryInt(bc.a)
			numB, _ := NewArbitraryInt(bc.b)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				numA.Add(numB)
			}
		})
	}
}

func BenchmarkMultiply(b *testing.B) {
	benchmarkCases := []struct {
		name string
		a, b string
	}{
		{"SmallNumbers", "1234", "5678"},
		{"LargeNumbers", "123456789", "987654321"},
	}

	for _, bc := range benchmarkCases {
		b.Run(bc.name, func(b *testing.B) {
			numA, _ := NewArbitraryInt(bc.a)
			numB, _ := NewArbitraryInt(bc.b)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				numA.Multiply(numB)
			}
		})
	}
}

func BenchmarkDivide(b *testing.B) {
	benchmarkCases := []struct {
		name string
		a, b string
	}{
		{"SmallNumbers", "12345", "67"},
		{"LargeNumbers", "123456789012345678901234567890", "987654321"},
	}

	for _, bc := range benchmarkCases {
		b.Run(bc.name, func(b *testing.B) {
			numA, _ := NewArbitraryInt(bc.a)
			numB, _ := NewArbitraryInt(bc.b)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				numA.Divide(numB)
			}
		})
	}
}

func TestArithmeticOperations(t *testing.T) {
	type testCase struct {
		a, b, expected string
	}

	testCases := struct {
		Addition       []testCase
		Subtraction    []testCase
		Multiplication []testCase
		Division       []struct {
			a, b, expectedQuot, expectedRem string
		}
	}{
		Addition: []testCase{
			{"0", "0", "0"},
			{"123", "456", "579"},
			{"999999", "1", "1000000"},
			{"-100", "50", "-50"},
		},
		Subtraction: []testCase{
			{"0", "0", "0"},
			{"500", "300", "200"},
			{"-100", "50", "-150"},
			{"1000", "1", "999"},
		},
		Multiplication: []testCase{
			{"0", "1000", "0"},
			{"10", "20", "200"},
			{"-5", "7", "-35"},
			{"999999", "999999", "999998000001"},
		},
		Division: []struct {
			a, b, expectedQuot, expectedRem string
		}{
			{"10", "3", "3", "1"},
			{"100", "10", "10", "0"},
			{"0", "5", "0", "0"},
		},
	}

	// Helper function to reduce code duplication
	runTest := func(t *testing.T, name string,
		testFunc func(a, b *ArbitraryInt) interface{},
		cases []testCase) {
		t.Run(name, func(t *testing.T) {
			for _, tc := range cases {
				numA, errA := NewArbitraryInt(tc.a)
				numB, errB := NewArbitraryInt(tc.b)

				if errA != nil || errB != nil {
					t.Fatalf("Failed to create ArbitraryInt: a=%v, b=%v", errA, errB)
				}

				var result string
				switch v := testFunc(numA, numB).(type) {
				case *ArbitraryInt:
					result = v.String()
				case struct {
					quotient, remainder *ArbitraryInt
				}:
					result = v.quotient.String()
				default:
					t.Fatalf("Unexpected return type")
				}

				if result != tc.expected {
					t.Errorf("%s op %s: Expected %s, got %s",
						tc.a, tc.b, tc.expected, result)
				}
			}
		})
	}

	runTest(t, "Addition", func(a, b *ArbitraryInt) interface{} {
		return a.Add(b)
	}, testCases.Addition)

	runTest(t, "Subtraction", func(a, b *ArbitraryInt) interface{} {
		return a.Subtract(b)
	}, testCases.Subtraction)

	runTest(t, "Multiplication", func(a, b *ArbitraryInt) interface{} {
		return a.Multiply(b)
	}, testCases.Multiplication)

	t.Run("Division", func(t *testing.T) {
		for _, tc := range testCases.Division {
			numA, errA := NewArbitraryInt(tc.a)
			numB, errB := NewArbitraryInt(tc.b)

			if errA != nil || errB != nil {
				t.Fatalf("Failed to create ArbitraryInt: a=%v, b=%v", errA, errB)
			}

			quotient, remainder, err := numA.Divide(numB)
			if err != nil {
				t.Fatalf("Unexpected division error: %v", err)
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
