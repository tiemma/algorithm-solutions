package main

import (
	"flag"
	"fmt"
	"math/big"
)

// https://projecteuler.net/problem=71

func main() {
	window := flag.Int64("window", 1_000_000, "The window size")
	flag.Parse()

	num, denum := int64(3), int64(7)
	upperFrac := big.NewRat(3, 7)
	lowerFrac := big.NewRat(2, 5)
	minRat := big.NewRat(1, 1)
	maxRat := big.NewRat(0, 1)
	result := big.NewRat(1, 1)

	zero := new(big.Rat)
	for {
		if denum > *window || num > *window {
			break
		}

		currRat := big.NewRat(num, denum)
		newMinRat := new(big.Rat).Sub(upperFrac, currRat)
		newMaxRat := new(big.Rat).Sub(currRat, lowerFrac)
		if newMinRat.Cmp(minRat) < 0 && newMaxRat.Cmp(maxRat) > 0 && newMinRat.Cmp(zero) > 0 && newMaxRat.Cmp(zero) > 0 {
			minRat = newMinRat
			maxRat = newMaxRat
			result = currRat
		}

		if big.NewRat(num, denum).Cmp(upperFrac) > 0 {
			denum += 1
		} else if big.NewRat(num, denum).Cmp(lowerFrac) < 0 {
			num += 1
		} else {
			num += 1
			denum += 1
		}
	}

	fmt.Println(result)
}
