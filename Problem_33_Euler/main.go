package main

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func main() {
	dp, np := 1, 1
	for c := 1; c <= 9; c++ {
		for d := 1; d < c; d++ {
			for n := 1; n < d; n++ {
				if ((n * 10) + c) * d == ((c * 10) + d) * n {
					np *= n
					dp *= d
					fmt.Println(c, n, d)
				}
			}
		}
	}

	fmt.Println(dp / gcd(np, dp))
}