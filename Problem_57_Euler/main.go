package main

import (
	"flag"
	"fmt"
	"math/big"
)

// https://projecteuler.net/problem=55

func main() {
	window := flag.Int("window", 1_000, "Range to find the longest lychrel number")
	flag.Parse()

	count := 0

	numerator, denominator := big.NewInt(3), big.NewInt(2)
	for i := 0; i < *window; i++ {
		oldNumerator := big.NewInt(1)
		oldNumerator.Mul(oldNumerator, numerator)

		oldDenominator := big.NewInt(1)
		oldDenominator.Mul(oldDenominator, denominator)

		numerator.Add(numerator, oldDenominator.Mul(oldDenominator, big.NewInt(2)))
		denominator.Add(denominator, oldNumerator)
		if len(numerator.String()) != len(denominator.String()) {
			fmt.Println(numerator, denominator)
			count += 1
		}
	}

	fmt.Println("-----------")
	fmt.Println(count)
}
