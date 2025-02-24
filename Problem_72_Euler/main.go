package main

import (
	"flag"
	"fmt"
	"math"
)

// https://projecteuler.net/problem=72

func isPrime(n int) bool {
	sqrt := int(math.Sqrt(float64(n))) + 1
	for i := 2; i < sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func eulerTotientMethod(window int) int {
	totalFractions := 0.0
	var primes []int

	for i := 2; i <= window; i++ {
		count := float64(i)
		if isPrime(i) {
			primes = append(primes, i)
			count -= 1
		} else {
			// https://en.wikipedia.org/wiki/Euler%27s_totient_function#Computing_Euler's_totient_function
			for _, p := range primes {
				if i%p == 0 {
					count *= 1 - (1 / float64(p))
				}
			}
		}

		totalFractions += count
	}

	return int(totalFractions)
}

func main() {
	window := flag.Int("window", 1_000_000, "The window size")
	flag.Parse()

	data := []int{0, 0}
	for i := 2; i < *window+1; i++ {
		data = append(data, i-1)
	}

	for i := 2; i < len(data); i++ {
		for j := i * 2; j < len(data); j += i {
			data[j] -= data[i]
		}
	}

	maxVal := 0
	for i, _ := range data {
		maxVal += data[i]
	}

	fmt.Println(maxVal)
}
