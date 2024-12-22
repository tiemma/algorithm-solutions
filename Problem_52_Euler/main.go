package main

import (
	"flag"
	"fmt"
	"math"
)

// https://projecteuler.net/problem=52

func breakNumber(n int) []int {
	result := []int{}
	for n > 0 {
		result = append([]int{n % 10}, result...)
		n /= 10
	}

	return result
}

func isEqual(a map[int]int, b map[int]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k1, v1 := range a {
		v2, ok := b[k1]
		if !ok || v2 != v1 {
			return false
		}
	}

	return true
}

func getDigitCounts(number int) map[int]int {
	digits := breakNumber(number)
	digitCounts := map[int]int{}
	for _, digit := range digits {
		digitCounts[digit]++
	}

	return digitCounts
}

func isPermutedMultiple(prevDigitCounts map[int]int, window int, number int) bool {
	for multiplier := 2; multiplier <= window; multiplier++ {
		if !isEqual(prevDigitCounts, getDigitCounts(multiplier*number)) {
			return false
		}
	}

	return true
}

func main() {
	window := flag.Int("window", 6, "Range of multiplier")
	flag.Parse()

	for i := 2; i < math.MaxInt64; i++ {
		prevDigitCounts := getDigitCounts(i)
		if isPermutedMultiple(prevDigitCounts, *window, i) {
			fmt.Println(i)
			return
		}
	}
}
