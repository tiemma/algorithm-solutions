package main

// https://projecteuler.net/problem=2


import (
	"fmt"
	"math"
	)

func BinetFormula(n float64) float64 {
	sqrt5 := math.Sqrt(5)
	phi := (1 + sqrt5) / 2
	ans := math.Round(math.Pow(phi, n) / sqrt5)
	fmt.Printf("Value at n: %f in fibonnaci sequence is: %f\n", n, ans)
	return ans
}

// O(n) time and space
// n - Number of even fibonacci occurrences
func FibonnaciSum() int {
	currValue := 0.0
	n := 3.0
	maxValue := 4* math.Pow(10, 6)
	for BinetFormula(n) < maxValue {
		currValue += BinetFormula(n)
		fmt.Println(currValue)
		// 1 2 3 5 8 13 21 34....
		// Following the sequence, we see that the index for even numbers occurs at steps of 3
		// 1 4 7 ....
		// So no need to bother with the other sequences, just increment the index by 3 and move to the next even one
		n += 3
	}
	return int(currValue)
}

func main(){
	fmt.Println(FibonnaciSum())
}
