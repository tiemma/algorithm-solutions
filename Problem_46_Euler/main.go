package main

import (
	"fmt"
	"math"
)

// https://projecteuler.net/problem=46

func isPrime(n int) bool {
	sqrt := int(math.Sqrt(float64(n))) + 1
	for i := 2; i < sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func isValidGoldbachConjecture(n int, primes []int) bool {
	for _, prime := range primes {
		if n < prime {
			break
		}

		number := math.Sqrt(0.5 * float64(n-prime))

		if math.Floor(number) == number {
			return true
		}
	}

	return false
}

func main() {
	primes := []int{1, 2}
	for i := 3; i < math.MaxInt64; i += 2 {
		if isPrime(i) {
			primes = append(primes, i)
		}

		if !isValidGoldbachConjecture(i, primes) {
			fmt.Println(i)
			return
		}
	}
}
