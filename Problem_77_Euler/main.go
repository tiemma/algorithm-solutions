package main

import (
	"flag"
	"fmt"
	"math"
)

// https://projecteuler.net/problem=77

func primePartitions(n int, primes []int, currentIndex int, memo map[string]int) int {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	memoKey := fmt.Sprintf("%d-%d", n, currentIndex)
	if val, found := memo[memoKey]; found {
		return val
	}

	totalPartitions := 0
	for i := currentIndex; i < len(primes); i++ {
		prime := primes[i]
		totalPartitions += primePartitions(n-prime, primes, i, memo)
	}

	memo[memoKey] = totalPartitions
	return totalPartitions
}

func isPrime(n int) bool {
	i := 2
	for i <= int(math.Sqrt(float64(n))) {
		if n%i == 0 {
			return false
		}
		i++
	}
	return true
}

func primePartitionsCount(n int, primes []int) int {
	memo := make(map[string]int)

	return primePartitions(n, primes, 0, memo)
}

func main() {
	window := flag.Int("window", 5000, "window size")
	flag.Parse()

	var primes []int
	for i := 2; i < *window; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	for i := 2; i < math.MaxInt64; i++ {
		result := primePartitionsCount(i, primes)
		fmt.Println(i, result)
		if result > *window {
			fmt.Println(i)
			return
		}
	}
}
