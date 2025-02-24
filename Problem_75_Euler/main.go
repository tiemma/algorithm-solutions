package main

// https://projecteuler.net/problem=75

import (
	"flag"
	"fmt"
	"math"
)

func bruteforce(window int) {
	perimeters := make([]int, window+1)
	for perimeter := 2; perimeter <= window; perimeter += 2 {
		if perimeters[perimeter] == 1 {
			perimeters[perimeter] = 0
			continue
		}
		for i := 2; i <= window; i++ {
			for j := i + 1; j < perimeter-i; j++ {
				k := perimeter - i - j
				if (i*i)+(j*j) == k*k {
					fmt.Println(perimeter, i, j, k)
					perimeters[perimeter] += 1
				}
			}
		}
	}

	count := 0
	for _, v := range perimeters {
		if v != 0 {
			count += 1 / v
		}
	}

	fmt.Println(count)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func main() {
	window := flag.Int("window", 1_500_000, "Window to search for the perimeter values")
	flag.Parse()

	//bruteforce(*window)
	fmt.Println("*******************************")
	//singleIntegerRightAngle(*window)
	eulersMethod(*window)
}

// This is so much faster than the method I derived for finding pythagorean triples
func eulersMethod(window int) {
	perimeters := map[int]int{}
	for m := 2; m <= int(math.Sqrt(float64(window/2)))+1; m++ {
		for n := 1; n <= m; n++ {
			if (m-n)%2 == 1 && gcd(m, n) == 1 {
				a := m*m - n*n
				b := 2 * m * n
				c := m*m + n*n
				perimeter := a + b + c
				//fmt.Println(perimeter, a, b, c)
				for j := perimeter; j <= window; j += perimeter {
					perimeters[j] += 1
				}
			}
		}
	}

	counts := 0
	for _, v := range perimeters {
		counts += 1 / v
	}

	fmt.Println(counts)
}

func singleIntegerRightAngle(window int) {

	// From my article: https://blog.bakman.build/blogs/unravelling-mathematical-threesomes.html
	// You can obtain a pythagorean triples using just one value "b" and some pythagorean distance "k"
	// So perimeter = a + b + c = (b^2 / k) + b
	// perimeter = (b^2 / k) + b
	// k = b ^ 2 / (perimeter - b)

	// To solve this, k must be an integer based on the equation above
	// In getting b, we can conclude that it is always bounded within
	// perimeter/2 as no pythagorean value can be greater than its perimeter halved.
	// If we consider that the longest side, the hypotenuse "c" is less than the perimeter halved:
	// (a + b + c) / 2 > c
	// a + b + c  > 2c
	// a + b > c
	// a + b > sqrt(a^2 + b^2) which will hold true as
	// (a + b)^ 2 > a^2 + b^2 by 2ab once you expand the factors

	// In deriving the solution for this problem, we can have permutations of the same tuple, this therefore checks
	// that the triple found in (a,b,c) has not been seen before.
	// We then increment the perimeter and adds the count at the end of the iterations from 0 to the window.
	perimeters := map[int]int{}
	for perimeter := 2; perimeter <= window; perimeter += 2 {
		seen := map[[3]int]bool{}
		for b := 2; b < perimeter/2; b += 2 {
			bSquared := b * b
			bsk := perimeter - b
			if bSquared%bsk != 0 {
				continue
			}

			k := bSquared / bsk
			if (bsk-k)%2 == 0 && bsk != k {
				a := (bsk - k) / 2
				c := perimeter - a - b
				triples := [3]int{a, b, c}
				if a > b {
					triples = [3]int{b, a, c}
				}
				if _, ok := seen[triples]; ok {
					continue
				}
				seen[triples] = true
				//fmt.Println(perimeter, a, b, c)
				perimeters[perimeter] += 1
			}
		}
	}

	fmt.Println("----------------")
	counts := 0
	for _, v := range perimeters {
		counts += 1 / v
	}

	fmt.Println(counts)
}
