package main

import (
	"flag"
	"fmt"
	"math/big"
)

// https://projecteuler.net/problem=48

func main() {
	maxNumPtr := flag.Int("number", 1000, "Integer for the max range")
	flag.Parse()

	sum := big.NewInt(0)
	for i := 1; i <= *maxNumPtr; i++ {
		fmt.Println(i)

		val := big.NewInt(int64(i))
		exp := val.Exp(val, val, big.NewInt(0))
		sum.Add(sum, exp)
	}

	sumAsString := sum.String()
	fmt.Println(sumAsString[len(sumAsString)-10:])
}
