package main

import (
	"fmt"
	"math"
)

// https://projecteuler.net/problem=44

func getPentagonalNumber(n int) int {
	return n * (3*n - 1) / 2
}

func main() {
	// In the series, it is quadratic and monotonically increasing
	// I assume a max range to search for the minimum distance between two pentagon numbers
	// After finding the first distance, any other distance greater than it confirms our proof of divergence and we can stop there
	values := map[int]struct{}{}
	minDistance := math.MaxInt
	for i := 1; i < math.MaxInt; i++ {
		number := getPentagonalNumber(i)
		values[number] = struct{}{}
		for j, _ := range values {
			if j > number/2 {
				continue
			}
			if _, ok := values[number-j]; ok {
				D := number - 2*j
				if _, ok := values[D]; ok {
					fmt.Println(i, number, j, number-j, D)
					if D < minDistance {
						minDistance = D
					} else {
						break
					}
				}
			}
		}

		if minDistance != math.MaxInt {
			break
		}
	}

	fmt.Println(minDistance)
}
