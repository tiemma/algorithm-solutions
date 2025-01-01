package main

import (
	"flag"
	"fmt"
	"math"
)

func isPrime(n int) bool {
	i := 2
	if n == 0 || n == 1 {
		return false
	}
	for i <= int(math.Sqrt(float64(n))) {
		if n%i == 0 {
			return false
		}
		i++
	}

	return true
}

func main() {
	window := flag.Int("window", 1_000_000, "The window size")
	flag.Parse()

	fmt.Println(eulerProductFormula(*window))
	return

	// https://projecteuler.net/overview=0069
	// This is so cracked, basically the max(n/totient(n)) is the max prime product to the number n
	n := 1
	for i := 2; n*i < *window; i++ {
		if isPrime(i) {
			fmt.Println(n)
			n *= i
		}
	}

	fmt.Println("------------------")
	fmt.Println(n)
}

func eulerProductFormula(window int) int {
	maxN, n := 0.0, 2
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

		phiN := float64(i) / count
		if phiN > maxN {
			//fmt.Println(phiN)
			maxN = phiN
			n = i
		}
	}

	return n
}
