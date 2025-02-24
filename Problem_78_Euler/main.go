package main

import (
	"flag"
	"fmt"
	"math"
	"math/big"
)

// https://projecteuler.net/problem=78

func getPentagonalNumber(n int) int {
	return n * (3*n - 1) / 2
}

func computePartitions(goal int, divisor int) (int, *big.Int) {
	partitions := []*big.Int{big.NewInt(1)}

	for n := 1; n <= goal; n++ {
		partitions = append(partitions, new(big.Int))
		for k := 1; k <= n; k++ {
			coefficient := int64(math.Pow(-1, float64(k+1)))
			for _, t := range []int{getPentagonalNumber(k), getPentagonalNumber(-k)} {
				if (n - t) >= 0 {
					partitions[n] = new(big.Int).Add(partitions[n], new(big.Int).Mul(big.NewInt(coefficient), partitions[n-t]))
				}
			}
		}

		currPartition := partitions[n]
		if new(big.Float).Quo(new(big.Float).SetInt(currPartition), big.NewFloat(float64(divisor))).IsInt() {
			return n, currPartition
		}
	}

	return 0, new(big.Int)
}

func main() {
	divisor := flag.Int("divisor", 1_000_000, "The window size")
	flag.Parse()

	fmt.Println(computePartitions(math.MaxInt64, *divisor))
}
