package main

import (
	"fmt"
	"math"
	"os"
)

// https://projecteuler.net/problem=45

func getQuadraticRoots(a, b, c float64) (float64, float64) {
	firstRoot := (-b + math.Sqrt(math.Pow(b, 2)-(4*a*c))) / (2 * a)
	secondRoot := (-b - math.Sqrt(math.Pow(b, 2)-(4*a*c))) / (2 * a)

	return firstRoot, secondRoot
}

func isValidRoot(ft, st float64) bool {
	for _, roots := range [][]float64{{ft, st}} {
		for _, root := range roots {
			if root > 0 && math.Floor(root) == root {
				return true
			}
		}
	}

	return false
}

func firstAndLastTwo(i int) (int, int) {
	return i / 100, i % 100
}

func sum(n []int) int {
	result := 0
	for _, n := range n {
		result += n
	}

	return result
}

func has(haystack []int, needle int) bool {
	for _, n := range haystack {
		if n == needle {
			return true
		}
	}

	return false
}

func filter(s string, haystack []string) []string {
	result := []string{}
	for _, hay := range haystack {
		if hay != s {
			result = append(result, hay)
		}
	}

	return result
}

func backtrack(shapes []string, digits []int, order []string, sets map[string][][]int) {
	if len(shapes) == 0 {
		first, _ := firstAndLastTwo(digits[0])
		_, lastD := firstAndLastTwo(digits[len(digits)-1])
		if lastD == first {
			fmt.Println("---------------")
			fmt.Println(digits, order)
			fmt.Println(sum(digits))
			// There is only one solution, so we can exit on the first version
			os.Exit(0)
		}
	} else {
		for j := 10; j < 100; j++ {
			for _, shape := range shapes {
				nums := sets[shape][j]
				newShapes := filter(shape, shapes)
				for _, num := range nums {
					newDigits := append(digits, num)
					newOrder := append(order, shape)
					if len(digits) == 0 {
						backtrack(newShapes, newDigits, newOrder, sets)
						continue
					}
					if has(digits, num) {
						continue
					}
					first, _ := firstAndLastTwo(num)
					_, lastD := firstAndLastTwo(digits[len(digits)-1])
					if lastD == first {
						backtrack(newShapes, newDigits, newOrder, sets)
					}
				}
			}
		}
	}
}

func main() {
	sets := map[string][][]int{}
	shapes := []string{"t", "s", "p", "hx", "hp", "o"}

	for _, set := range shapes {
		sets[set] = make([][]int, 100)
	}
	for k := 1000; k < 10000; k++ {
		firstTwo, lastTwo := firstAndLastTwo(k)
		i := float64(k)
		if firstTwo < 10 && lastTwo < 10 {
			continue
		}

		ft, st := getQuadraticRoots(0.5, 0.5, -i)
		if isValidRoot(ft, st) {
			sets["t"][firstTwo] = append(sets["t"][firstTwo], k)
		}

		fs, ss := getQuadraticRoots(1, 0, -i)
		if isValidRoot(fs, ss) {
			sets["s"][firstTwo] = append(sets["s"][firstTwo], k)
		}

		fp, sp := getQuadraticRoots(1.5, -0.5, -i)
		if isValidRoot(fp, sp) {
			sets["p"][firstTwo] = append(sets["p"][firstTwo], k)
		}

		fh, sh := getQuadraticRoots(2, -1, -i)
		if isValidRoot(fh, sh) {
			sets["hx"][firstTwo] = append(sets["hx"][firstTwo], k)
		}

		fhp, shp := getQuadraticRoots(2.5, -1.5, -i)
		if isValidRoot(fhp, shp) {
			sets["hp"][firstTwo] = append(sets["hp"][firstTwo], k)
		}

		fo, so := getQuadraticRoots(3, -2, -i)
		if isValidRoot(fo, so) {
			sets["o"][firstTwo] = append(sets["o"][firstTwo], k)
		}
	}

	backtrack(shapes, []int{}, []string{}, sets)
}
