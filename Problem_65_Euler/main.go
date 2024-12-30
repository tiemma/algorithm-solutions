package main

import (
	"flag"
	"fmt"
	"math"
	"math/big"
)

func main() {
	window := flag.Int("window", 10, "The window size")
	flag.Parse()

	zero := big.NewInt(0)
	ten := big.NewInt(10)
	n, m := []*big.Int{big.NewInt(1), big.NewInt(1)}, []*big.Int{big.NewInt(0), big.NewInt(1)}
	for i := 1; i < math.MaxInt64; i++ {
		index := i - 1
		multiplier := big.NewInt(1)
		fmt.Println(index, n[0], m[0])

		if index%3 == 0 {
			multiplier.Mul(multiplier, big.NewInt(2)).Mul(multiplier, big.NewInt(int64(index))).Div(multiplier, big.NewInt(3))
		}

		a0, a1 := n[0], n[1]
		newN := big.NewInt(1)
		n = []*big.Int{a1, newN.Mul(a1, multiplier).Add(a0, newN)}

		b0, b1 := m[0], m[1]
		newM := big.NewInt(1)
		m = []*big.Int{b1, newM.Mul(b1, multiplier).Add(b0, newM)}

		if index == *window+1 {
			break
		}
	}

	var result int64
	start := n[0]
	for start.Cmp(zero) > 0 {
		result += big.NewInt(1).Mod(start, ten).Int64()
		start.Div(start, ten)
	}
	fmt.Println(result)
}
