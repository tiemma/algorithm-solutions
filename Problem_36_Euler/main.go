package main

// https://projecteuler.net/problem=36

import (
	"flag"
	"fmt"
	"strconv"
)

func isPalindrome(numStr string) bool {
	return numStr == Reverse(numStr)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	window := flag.Float64("window", 1_000_000, "Window to search for circular primes")
	flag.Parse()

	result := 0
	for i := 1; i < int(*window); i++ {
		if isPalindrome(strconv.Itoa(i)) && isPalindrome(strconv.FormatInt(int64(i), 2)) {
			fmt.Println(i)
			result += i
		}
	}

	fmt.Println("-----------")
	fmt.Println(result)
}

