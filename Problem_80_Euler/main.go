package main

import (
	"flag"
	"fmt"
	"math/big"
	"strings"
)

// https://projecteuler.net/problem=80

// Use main.py as golang has precision issues
func main() {
	window := flag.Int("window", 100, "The window size")
	flag.Parse()

	prec := uint(200)
	total := 0
	for i := 0; i <= *window; i++ {
		val := new(big.Float).Sqrt(big.NewFloat(float64(i)).SetPrec(100)).SetPrec(prec)
		newVal := fmt.Sprintf("%.100f", val)
		dValue := strings.Split(newVal, ".")[1]
		fmt.Println(dValue)
		for _, char := range strings.TrimSpace(dValue) {
			fmt.Println(string(char), int(char-'0'))
			total += int(char - '0')
		}
	}

	fmt.Println(total)
}
