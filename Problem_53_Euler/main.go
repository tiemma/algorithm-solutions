package main

import (
	"fmt"
	"math/big"
)

// https://projecteuler.net/problem=53

func factorial(n int) *big.Int {
	result := big.NewInt(int64(1))

	if n == 0 {
		return result
	}
	for i := 1; i <= n; i++ {
		result = result.Mul(result, big.NewInt(int64(i)))
	}

	return result
}

func main() {
	oneMillion := big.NewInt(1_000_000)
	count := 0
	factorialMap := make(map[int]*big.Int)
	for i := 0; i <= 100; i++ {
		factorialMap[i] = factorial(i)
	}

	for n := 1; n <= 100; n++ {
		for r := 1; r < n; r++ {
			nf := big.NewInt(1).Mul(big.NewInt(1), factorialMap[n])
			rf := big.NewInt(1).Mul(big.NewInt(1), factorialMap[r])
			nmrf := big.NewInt(1).Mul(big.NewInt(1), factorialMap[n-r])
			if nf.Div(nf, rf.Mul(rf, nmrf)).Cmp(oneMillion) > 0 {
				count += 1
			}
		}
	}

	fmt.Println(count)
}
