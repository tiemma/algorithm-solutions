package main

import (
	"flag"
	"fmt"
	"math"
)

// https://projecteuler.net/problem=73

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func main() {
	window := flag.Int("window", 12_000, "The window size")
	flag.Parse()

	totalSum := 0

	// 2/5
	// 3/7 3/8
	// 4/9 4/10 4/11
	// 5/11 5/12 5/13 5/14
	// 6/13 6/14 6/15 6/16 6/17
	// 7/15 7/16 7/17 7/18 7/19 7/20

	for i := 1; i <= *window/2; i++ {
		for j := 2*i + 1; j < 3*i; j += 1 {
			if j > *window {
				break
			}
			//fmt.Println(i, j)
			if gcd(i, j) == 1 {
				totalSum += 1
			}
		}
	}

	fmt.Println(totalSum)
}

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

func bruteforce(window *int) {

	totalSum := 0

	var primes []int
	numReductions := map[int]int{}
	denumReductions := map[int]int{}

	for j := 2; j <= *window; j++ {
		if isPrime(j) {
			primes = append(primes, j)
		}
	}

	// 2/5
	// 3/7 3/8
	// 4/9 4/10 4/11
	// 5/11 5/12 5/13 5/14
	// 6/13 6/14 6/15 6/16 6/17
	// 7/15

	for i := 1; i <= *window/2; i++ {
		for j := 2*i + 1; j < 3*i; j += 1 {
			if j > *window {
				break
			}
			fmt.Println("+", i, j)
			totalSum += 1
			for _, prime := range primes {
				if prime > i {
					break
				}
				if i%prime == 0 && j%prime == 0 {
					numReductions[i] += 1
					denumReductions[j] += 1
					fmt.Println("-", i, j)
					totalSum -= 1
					break
				}
			}
		}
	}

	//n(1 + 1.25(n-10))
	fmt.Println(numReductions)
	fmt.Println(denumReductions)
	fmt.Println(totalSum)
}
