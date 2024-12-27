package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

// https://projecteuler.net/problem=51

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

func getValue(digits []int) int {
	value := 0
	power := 0
	for i := 0; i < len(digits); i++ {
		number := float64(digits[len(digits)-i-1])
		value += int(math.Pow(10, float64(power)) * number)
		power += int(math.Floor(math.Log10(number)) + 1)
	}

	return value
}

func isPrimePair(a, b int) bool {
	for _, number := range [][]int{{a, b}, {b, a}} {
		if !isPrime(getValue(number)) {
			return false
		}
	}

	return true
}

func sum(n []int) int {
	result := 0
	for _, n := range n {
		result += n
	}

	return result
}

func getPermutations(arr, rest []int, n int, result [][]int) [][]int {
	if n == 0 {
		return [][]int{arr}
	} else {
		for i, k := range rest {
			result = append(result, getPermutations(append([]int{k}, arr...), rest[i+1:], n-1, [][]int{})...)
		}
	}

	return result
}

func getPermutationsChan(arr, rest []int, n int, wg chan []int) {
	if n == 0 {
		wg <- arr
	} else {
		for i, k := range rest {
			getPermutationsChan(append([]int{k}, arr...), rest[i+1:], n-1, wg)
		}
	}
}

func processResult(wg chan []int, i int) {
	for result := range wg {
		//fmt.Println(result, i)
		permutations := getPermutations([]int{}, result, 2, [][]int{})
		valid := true
		for _, perm := range permutations {
			if !isPrimePair(perm[0], perm[1]) {
				valid = false
				break
			}
		}

		if valid {
			fmt.Println("------------")
			fmt.Println(sum(result), result)
			close(wg)

			os.Exit(0)
		}
	}
}

func main() {
	window := flag.Int("window", 4, "Number of prime factors to find replacements matching")
	flag.Parse()

	power := 0
	var primes []int
	for power < *window {
		for i := int(math.Pow(10, float64(power-1))); i < int(math.Pow(10, float64(power))); i++ {
			if isPrime(i) {
				primes = append(primes, i)
			}
		}

		wg := make(chan []int)
		for i := 0; i < 1000; i++ {
			go processResult(wg, i)
		}

		getPermutationsChan([]int{}, primes, *window, wg)

		power += 1
	}
}
