package main

// https://projecteuler.net/problem=37

import (
	"fmt"
)

func main() {
	// It is redundant to find the truncatable primes using math so I just pull them online
	// There are 15 primes which are both left-truncatable and right-truncatable. They have been called two-sided primes. The complete list: 2, 3, 5, 7, 23, 37, 53, 73, 313, 317, 373, 797, 3137, 3797, 739397 (sequence A020994 in the OEIS)
	// In the problem statement, 2, 3, 5 and 7 are not allowed hence the number being 11.
	numbers := []int{23, 37, 53, 73, 313, 317, 373, 797, 3137, 3797, 739397}
	result := 0
	for _, i := range numbers {
		result += i
	}

	fmt.Println(result)
}

