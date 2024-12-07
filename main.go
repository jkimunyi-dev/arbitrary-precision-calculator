package main

import (
	"arbitrary-precision-calculator/internal"
	"fmt"
	"log"
)

func main() {
	// Example usage of ArbitraryInt
	num, err := internal.NewArbitraryInt("12345678901234567890")
	if err != nil {
		log.Fatalf("Error creating number: %v", err)
	}

	fmt.Println("Large Number:", num)
}
