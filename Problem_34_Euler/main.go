package main

// https://projecteuler.net/problem=34

import (
	"fmt"
)


func FactorialSum(prev map[int]int, num int) int {
	result := 0
	for num > 0 {
		remainder := num - (10 * (num/10))
		result += prev[remainder]
		num = num/10
	}

	return result
}


func main() {
	prevFactorial := map[int]int{0: 1, 1: 1}
	for i := 2; i < 10; i++ {
		prevFactorial[i] = prevFactorial[i-1] * i
	}
	fmt.Println(prevFactorial)


	i := 3
	result := 0
	//map[0:1 1:1 2:2 3:6 4:24 5:120 6:720 7:5040 8:40320 9:362880]
	// To get the largest factorial that can be a sum, it cannot be bigger than the largest factorial starting
	// the sum which is 9!. Anything beyond that is impossible.
	for i < prevFactorial[9] {
		factorialSum := FactorialSum(prevFactorial, i)
		if i == factorialSum {
			fmt.Println(i, factorialSum)
			result += i
		}
		i++
	}

	fmt.Println(result)
}

