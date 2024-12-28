package main

import (
	"flag"
	"fmt"
	"math"
)

// https://projecteuler.net/problem=57

func gcd(a, b int) int {
	a = int(math.Abs(float64(a)))
	b = int(math.Abs(float64(b)))
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func getDenominator(number, lastRoot int) int {
	return number - int(math.Pow(float64(lastRoot), 2))
}

func getNumerator(denominator, numerator int) int {
	return denominator * numerator
}

type values struct {
	numerator   int
	denominator int
	roots       []int
}

func backtrack(n, d, number int) []int {
	var stack []values
	stack = append(stack, values{n, d, []int{n}})
	for len(stack) > 0 {
		stackValue := stack[0]
		stack = stack[1:]

		roots := stackValue.roots
		numerator := stackValue.numerator
		denominator := stackValue.denominator

		newDenominator := getDenominator(number, numerator)
		newNumerator := getNumerator(denominator, numerator)

		divisor := gcd(denominator, newDenominator)
		newNumerator /= divisor
		newDenominator /= divisor
		if newNumerator <= 0 || newDenominator <= 0 {
			continue
		}
		if len(roots) > 1 && numerator == n && denominator == d {
			return roots
		}

		for j := 1; j <= (n+1)*2; j++ {
			adjustedNumerator := (j * newDenominator) - newNumerator
			if adjustedNumerator <= 0 {
				continue
			}
			stack = append(stack, values{adjustedNumerator, newDenominator, append([]int{j}, roots...)})
		}
	}

	return []int{n}
}

// Obtained from https://en.wikipedia.org/wiki/Periodic_continued_fraction#Canonical_form_and_repetend
func getCount(n int) int {
	m := 0
	d := 1
	a0 := int(math.Sqrt(float64(n)))
	a := a0
	count := 1
	for {
		m = d*a - m
		denominator := getDenominator(n, m)
		if denominator == 0 {
			return 0
		}
		d = denominator / d
		a = (a0 + m) / d
		//fmt.Println(d, m, a)
		if a == 2*a0 {
			//fmt.Println(n, count)
			break
		}
		count += 1
	}

	return count
}

func main() {
	window := flag.Int("window", 10_000, "Range to find the longest lychrel number")
	flag.Parse()

	count := 0
	for i := 2; i <= *window; i++ {
		//numSqrt := math.Sqrt(float64(i))
		//baseRoot := int(numSqrt)
		//numerator, denominator := baseRoot, 1
		//resp := backtrack(numerator, denominator, i)
		//count += (len(resp) + 1) % 2
		count += getCount(i) % 2
		//fmt.Println(i, count)
		//return
	}
	fmt.Println(count)
}
