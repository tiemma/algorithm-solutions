package main

import (
	"fmt"
	"math"
)

// https://projecteuler.net/problem=47

func isPrime(n int) bool {
	sqrt := int(math.Sqrt(float64(n))) + 1
	for i := 2; i < sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func hasNConsecutivePrimes(number int, primes []int, n int) bool {
	count := 0
	for _, prime := range primes {
		if number%prime == 0 {
			count += 1
		}
	}

	return count == n
}

func main() {
	primes := []int{2}
	count := 0
	distinctFactors := 4
	for i := 3; i < math.MaxInt64; i += 1 {
		if isPrime(i) {
			primes = append(primes, i)
		}

		if hasNConsecutivePrimes(i, primes, distinctFactors) {
			count += 1
		} else {
			count = 0
		}

		if count == distinctFactors {
			fmt.Println(i - distinctFactors + 1)
			break
		}
	}
}
