package main

// https://projecteuler.net/problem=38

import (
	"flag"
	"fmt"
	"math"
)




func main() {
	window := flag.Float64("window", 1_000, "Window to search for the perimeter values")
	flag.Parse()

	// From my article: https://blog.bakman.build/blogs/unravelling-mathematical-threesomes.html
	// You can obtain a pythagorean triples using just one value "b" and some pythagorean distance "k"
	// So perimeter = a + b + c = (b^2 / k) + b
	// perimeter = (b^2 / k) + b
	// k = b ^ 2 / (perimeter - b)

	// To solve this, k must be an integer based on the equation above
	// In getting b, we can conclude that it is always bounded within
	// perimeter/2 as no pythagorean value can be greater than its perimeter halved.
	// If we consider that the longest side, the hypotenus "c" is less than the perimeter halved:
	// (a + b + c) / 2 > c
	// a + b + c  > 2c
	// a + b > c
	// a + b > sqrt(a^2 + b^2) which will hold true as
	// (a + b)^ 2 > a^2 + b^2 by 2ab once you expand the factors

	// The solution is maximised when the greater number of whole values of k that give whole values of a, b and c are obtained
	// This is therefore a search for the number of values from 0 to perimeter / 2 that satisfies those constraints
	maxPerimeter, maxCount := math.MinInt64, 0
	for perimeter := 2; perimeter < int(*window); perimeter++ {
		count := 0
		for b := 2; b < perimeter / 2; b++ {
			if int(math.Pow(float64(b), 2)) % int(perimeter - b) == 0 {
				k := math.Pow(float64(b), 2) / float64(perimeter - b)
				bsquared := math.Pow(float64(b), 2)
				c := 0.5 * (bsquared / k + k)
				a := 0.5 * (bsquared / k - k)
				// Fractional results are not needed here
				if c > math.Floor(c) || a > math.Floor(a) {
					continue
				}
				count += 1
			}
		}

		if count > maxCount {
			maxPerimeter = perimeter
			maxCount = count
			fmt.Println(maxPerimeter, maxCount)
		}
	}

	fmt.Println("----------------")
	// We will have a and b as individual values for the result hence we get double the actual number of values
	// It doesn't matter anyways as the answer does not require unique entries for the pythagorean triples
	fmt.Println(maxPerimeter, maxCount)
}

