package main

import (
	"flag"
	"fmt"
	"math/big"
)

// https://projecteuler.net/problem=55

func reverseNumber(n *big.Int) *big.Int {
	reverse := big.NewInt(0)
	ten := big.NewInt(10)

	newN := big.NewInt(0)
	newN.Add(newN, n)

	for newN.Cmp(big.NewInt(0)) > 0 {
		lastDigit := big.NewInt(1).Mod(newN, ten)
		reverse.Mul(reverse, ten).Add(reverse, lastDigit)
		newN = newN.Div(newN, ten)
	}

	return reverse
}

func main() {
	window := flag.Int("window", 10_000, "Range to find the longest lychrel number")
	flag.Parse()

	count := 0

	for i := 0; i < *window; i++ {
		number := big.NewInt(int64(i))
		maxIterations := 50

		for maxIterations > 0 {
			reversedNumber := reverseNumber(number)
			number.Add(number, reversedNumber)
			maxIterations--
			if number.Cmp(reverseNumber(number)) == 0 {
				break
			}
		}

		if maxIterations == 0 {
			fmt.Println(i)
			count += 1
		}
	}

	fmt.Println("---------------")
	fmt.Println(count)
}
