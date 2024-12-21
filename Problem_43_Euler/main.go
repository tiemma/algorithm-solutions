package main

import (
	"fmt"
	"math"
)

// https://projecteuler.net/problem=43

// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func getPermutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func getValue(digits []int) int {
	value := 0
	for i := 0; i < len(digits); i++ {
		value += int(math.Pow(10, float64(i))) * digits[len(digits)-i-1]
	}

	return value
}

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
	digits := []int{}
	for i := 0; i < 10; i++ {
		digits = append(digits, i)
	}

	maxPandigitalPrime := 0

	primes := map[int]int{}
	for i := 2; i < math.MaxInt64; i++ {
		if len(primes) > 6 {
			break
		}

		if isPrime(i) {
			primes[len(primes)+1] = i
		}
	}

	for _, perm := range getPermutations(digits) {
		number := getValue(perm)
		isDivisible := true

		for i := 1; i < len(digits)-2; i++ {
			newDigits := perm[i : i+3]
			if getValue(newDigits)%primes[i] != 0 {
				isDivisible = false
				break
			}
		}

		if isDivisible {
			fmt.Println(number)
			maxPandigitalPrime += number
		}

	}

	fmt.Println("---------------")
	fmt.Println(maxPandigitalPrime)
}
