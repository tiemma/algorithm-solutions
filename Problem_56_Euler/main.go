package main

import (
	"flag"
	"fmt"
	"math"
	"math/big"
)

// https://projecteuler.net/problem=55

func main() {
	window := flag.Int("window", 100, "Range to find the longest lychrel number")
	flag.Parse()

	ten := big.NewInt(10)
	maxSum := math.MinInt64

	for i := 1; i < *window; i++ {
		for j := 1; j < *window; j++ {
			a := big.NewInt(int64(i))
			b := big.NewInt(int64(j))
			c := a.Exp(a, b, nil)
			sum := 0

			for c.Cmp(big.NewInt(0)) != 0 {
				sum += int(big.NewInt(1).Mod(c, ten).Int64())
				c.Div(c, ten)
			}

			if sum > maxSum {
				fmt.Println(i, j, sum)
				maxSum = sum
			}
		}
	}

	fmt.Println("------------------")
	fmt.Println(maxSum)

}
