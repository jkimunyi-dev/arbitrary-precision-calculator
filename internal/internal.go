package internal

// ArbitraryInt represents an arbitrary precision integer
type ArbitraryInt struct {
	// Digits stored from least significant to most significant
	// This allows easier manipulation of large numbers
	digits []int
	// Sign of the number (true for negative, false for positive)
	negative bool
}
