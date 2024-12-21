package main

// https://projecteuler.net/problem=35

import (
	"flag"
	"fmt"
	"math"
	"strconv"
)


func isPrime(n int) bool {
	sqrt := int(math.Sqrt(float64(n))) + 1
	for i := 2; i < sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func rotateDigit(n int) int {
	remainder := n - (10 * (n/10))
	power := float64(len(strconv.Itoa(n))) - 1

	return (remainder * int(math.Pow(10, power))) + (n/10)
}

func isCircularPrime(n int) bool {
	newDigit := rotateDigit(n)

	for n != newDigit {
		if !isPrime(newDigit) {
			return false
		}
		newDigit = rotateDigit(newDigit)
	}

	return true
}


func main() {
	window := flag.Float64("window", 1_000_000, "Window to search for circular primes")
	flag.Parse()

	count := 0
	for i := 2; i < int(*window); i++ {
		if isPrime(i) && isCircularPrime(i) {
			fmt.Println(i)
			count += 1
		}
	}

	fmt.Println("-----------")
	fmt.Println(count)
}

