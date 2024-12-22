package main

import (
	"flag"
	"fmt"
	"math"
	"math/bits"
	"strconv"
)

// https://projecteuler.net/problem=51

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

func getValue(digits []int) int {
	value := 0
	for i := 0; i < len(digits); i++ {
		value += int(math.Pow(10, float64(i))) * digits[len(digits)-i-1]
	}

	return value
}

func breakNumber(n int) []int {
	result := []int{}
	for n > 0 {
		result = append([]int{n % 10}, result...)
		n /= 10
	}

	return result
}

// Combinations returns combinations of n elements for a given generic array.
// For n < 1, it equals to All and returns all combinations.
func Combinations[T any](set []T, n int) (subsets [][]T) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset []T

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}

func generateCombinations(numbers []int, combinations int) [][]int {
	return Combinations(numbers, combinations)
}

func getReplacementIndexes(number []int, digit int) []int {
	replacementIndexes := []int{}
	for i := 0; i < len(number); i++ {
		if number[i] == digit {
			replacementIndexes = append(replacementIndexes, i)
		}
	}

	return replacementIndexes
}

func getNumLength(num int) int {
	return len(strconv.Itoa(num))
}

func getPrimesFromReplacementIndex(number []int, replacements []int, digit int) []int {
	primes := []int{}
	for i := 0; i < 10; i++ {
		newDigits := make([]int, len(number))
		copy(newDigits, number)
		for _, v := range replacements {
			newDigits[v] = i
		}

		newNumber := getValue(newDigits)
		if getNumLength(newNumber) != len(number) {
			continue
		}
		if isPrime(newNumber) {
			primes = append(primes, newNumber)
		}
	}

	return primes
}

func main() {
	window := flag.Int("window", 6, "Number of prime factors to find replacements matching")
	flag.Parse()

	for i := 10; i < math.MaxInt64; i++ {
		if isPrime(i) {
			digits := breakNumber(i)
			digitCounts := map[int]int{}
			for _, digit := range digits {
				digitCounts[digit]++
			}

			for k, v := range digitCounts {
				for j := 1; j <= v; j++ {
					indexes := getReplacementIndexes(digits, k)
					for _, c := range generateCombinations(indexes, j) {
						primes := getPrimesFromReplacementIndex(digits, c, k)
						if len(primes) == *window {
							fmt.Println(primes)
							return
						}
					}
				}
			}
		}

	}
}
