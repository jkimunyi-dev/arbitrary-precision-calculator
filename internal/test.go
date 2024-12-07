package internal

import "testing"

func TestNewArbitraryInt(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		isError  bool
	}{
		{"12345", "12345", false},
		{"-12345", "-12345", false},
		{"+12345", "12345", false},
		{"0", "0", false},
		{"invalid", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			num, err := NewArbitraryInt(tc.input)

			if tc.isError {
				if err == nil {
					t.Errorf("Expected an error for input %s, but got none", tc.input)
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error for input %s: %v", tc.input, err)
				return
			}

			if num.String() != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, num.String())
			}
		})
	}
}

func TestCompare(t *testing.T) {
	testCases := []struct {
		a, b     string
		expected int
	}{
		{"100", "99", 1},
		{"99", "100", -1},
		{"100", "100", 0},
		{"-100", "99", -1},
		{"100", "-99", 1},
	}

	for _, tc := range testCases {
		t.Run(tc.a+" vs "+tc.b, func(t *testing.T) {
			numA, _ := NewArbitraryInt(tc.a)
			numB, _ := NewArbitraryInt(tc.b)

			result := numA.Compare(numB)
			if result != tc.expected {
				t.Errorf("Expected comparison result %d, got %d", tc.expected, result)
			}
		})
	}
}
