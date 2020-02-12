package main


// https://projecteuler.net/problem=3

import (
	"fmt"
	"math"
)

func DivMod(n float64, d float64) (float64, int) {
	return n / d, int(n) % int(d)
}

// O(n) time and O(1) space
func PrimeCheck(n float64) float64 {
	largestSqrt := math.Round(math.Sqrt(n))
	ans := 0.0
	for i := 2.0; i < largestSqrt; i++ {
		temp, rem := DivMod(n, i)
		if rem == 0 {
			n = temp
			ans = math.Max(ans, i)
		}
	}
	ans = math.Max(ans, n)

	return ans
}


func main(){
	fmt.Println(PrimeCheck(600851475143))
}