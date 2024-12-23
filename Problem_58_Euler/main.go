package main

import (
	"flag"
	"fmt"
	"math"
)

// https://projecteuler.net/problem=58

func isPrime(n int) bool {
	i := 2
	for i <= int(math.Sqrt(float64(n))) {
		if n%i == 0 {
			return false
		}
		i++
	}
	return true
}

func main() {
	window := flag.Float64("window", 0.1, "Ratio of prime diagonals to stop at")
	flag.Parse()

	count := 0
	for i := 3; i < math.MaxInt64; i += 2 {
		leftBottomDiagonal := i * i
		rightTopDiagonal := leftBottomDiagonal - (i-1)*2
		rightBottomDiagonal := leftBottomDiagonal - (i - 1)
		leftTopDiagonal := ((i - 2) * (i - 2)) + (i - 1)
		for _, number := range []int{leftBottomDiagonal, leftTopDiagonal, rightTopDiagonal, rightBottomDiagonal} {
			if isPrime(number) {
				count += 1
			}
		}

		numberOfPrimeDiagonals := float64(2*i - 1)
		if float64(count)/numberOfPrimeDiagonals < *window {
			fmt.Println(i, count, numberOfPrimeDiagonals)
			return
		}
	}
}
