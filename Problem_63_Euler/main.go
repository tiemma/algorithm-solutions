package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	result := 0
	for i := 1; i < 10; i++ {
		count := 0
		number := big.NewInt(int64(i))
		for n := 1; n <= math.MaxInt64; n++ {
			if len(number.String()) != n {
				fmt.Println(i, n-1)
				break
			}
			number.Mul(number, big.NewInt(int64(i)))
			count += 1
		}
		result += count
	}

	fmt.Println("-----------")
	fmt.Println(result)
}
