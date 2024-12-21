package main

import (
	"fmt"
	"math"
)

// https://projecteuler.net/problem=45

func getQuadraticRoots(a, b, c float64) (float64, float64) {
	firstRoot := (-b + math.Sqrt(math.Pow(b, 2)-(4*a*float64(c)))) / (2 * a)
	secondRoot := (-b - math.Sqrt(math.Pow(b, 2)-(4*a*float64(c)))) / (2 * a)

	return firstRoot, secondRoot
}

func isValidRoot(ft, st, fp, sp, fh, sh float64) bool {
	for _, roots := range [][]float64{{ft, st}, {fp, sp}, {fh, sh}} {
		isValid := 0
		for _, root := range roots {
			if root > 0 && math.Floor(root) == root {
				isValid += 1
			}
		}

		if isValid == 0 {
			return false
		}
	}

	return true
}

func hexagonalNumber(c int) int {
	return c * (2*c - 1)
}

func main() {
	for k := 144; k < math.MaxInt; k++ {
		i := float64(hexagonalNumber(k))
		ft, st := getQuadraticRoots(0.5, 0.5, -i)
		fp, sp := getQuadraticRoots(1.5, -0.5, -i)
		fh, sh := getQuadraticRoots(2, -1, -i)

		if isValidRoot(ft, st, fp, sp, fh, sh) {
			fmt.Println(ft, st, fp, sp, fh, sh)
			fmt.Println(k, int64(i))

			return
		}
	}
}
