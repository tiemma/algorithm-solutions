package main

import (
	"flag"
	"fmt"
)

// https://projecteuler.net/problem=74

func breakNumber(n int) []int {
	result := []int{}
	for n > 0 {
		result = append([]int{n % 10}, result...)
		n /= 10
	}

	return result
}

func getFactorialSum(n int, factorials map[int]int) int {
	total := 0
	for _, j := range breakNumber(n) {
		total += factorials[j]
	}

	return total
}

func main() {
	window := flag.Int("window", 1_000_000, "The window size")
	nonRepeatingTerms := flag.Int("terms", 60, "Number of non-repeating terms to find")
	flag.Parse()

	run(*window, *nonRepeatingTerms)
}

func run(window int, nonRepeatingTerms int) {
	count := 0
	states := map[int]int{}
	factorials := map[int]int{0: 1}
	for i := 1; i < 10; i++ {
		factorials[i] = factorials[i-1] * i
	}

	results := map[int]int{}
	for i := 2; i < window; i++ {
		node := i
		sequence := map[int]int{}

		for {
			_, ok := sequence[node]
			_, ok2 := results[node]
			if ok || ok2 {
				fmt.Println(i, len(sequence))
				results[i] = len(sequence) + results[node]
				if results[i] == nonRepeatingTerms {
					count += 1
				}

				break
			}

			newNode, ok := states[node]
			if !ok {
				newNode = getFactorialSum(node, factorials)
				states[node] = newNode
			}

			sequence[node] = newNode
			node = newNode
		}
	}

	fmt.Println(count)
}
