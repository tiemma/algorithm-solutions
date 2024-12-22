package main

import (
	"flag"
	"fmt"
	"math"
)

// https://projecteuler.net/problem=50

func isPrime(n int) bool {
	sqrt := int(math.Sqrt(float64(n))) + 1
	for i := 2; i < sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	window := flag.Int("window", 1_000_000, "Max number to sum to")
	flag.Parse()

	//terms := 0
	primes := []int{}
	primesSum := []int{0}
	for i := 2; i < *window; i++ {
		if isPrime(i) {
			primes = append(primes, i)
			primesSum = append(primesSum, i+primesSum[len(primesSum)-1])
		}
	}

	maxSum := math.MinInt
	maxI, maxJ := 0, 0
	terms := 0

	for i := 0; i < len(primes); i++ {
		for j := 0; j < len(primesSum); j++ {
			if primesSum[j] == 0 {
				continue
			}

			if primesSum[j] > 0 && primesSum[j] < *window && isPrime(primesSum[j]) && j-i > terms {
				fmt.Println(i, j, primesSum[j])
				maxI = i
				maxJ = j
				terms = j - i
				maxSum = primesSum[j]
			}

			primesSum[j] -= primes[i]
		}
	}

	fmt.Println("------------------")
	for i := maxI; i < maxJ; i++ {
		fmt.Print(primes[i], " ")
	}
	fmt.Println()
	fmt.Println(maxJ-maxI, maxSum)
}
