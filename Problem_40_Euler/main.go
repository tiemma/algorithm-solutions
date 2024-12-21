package main

import (
	"flag"
	"fmt"
	"math"
)

// https://projecteuler.net/problem=40

func getValueAndRemainder(x int, y int) (int, int) {
	return x / y, x % y
}

func digit(num, place int) int {
	fmt.Println(num, place)
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}

func getDigitLengthAtIndex(idx int, numLength map[int]int) int {
	for i := 0; i < math.MaxInt64; i++ {
		if numLength[i] > idx {
			return i
		}
	}

	return -1
}

func getDigitAtIndex(idx int, numLength map[int]int) int {
	k := getDigitLengthAtIndex(idx, numLength)
	adjustedStart := numLength[k-1] + 1
	value, remainder := getValueAndRemainder(idx-adjustedStart, k)
	fmt.Println(value, remainder)

	return digit(int(math.Pow(10, float64(k-1)))+value, k-remainder)
}

func main() {
	maxPower := flag.Float64("window", 1_000_000, "Window to generate powers of 10 for digits")
	flag.Parse()

	numLength := map[int]int{}
	numLength[0] = 0
	for i := 1; i <= int(math.Log10(*maxPower)); i++ {
		numLength[i] = numLength[i-1] + int(math.Pow(10, float64(i))-math.Pow(10, float64(i-1)))*i
	}
	fmt.Println(numLength)

	result := 1
	for i := 0; i <= int(math.Log10(*maxPower)); i++ {
		idx := int(math.Pow(10, float64(i)))
		number := getDigitAtIndex(idx, numLength)
		result = result * number
		fmt.Println(idx, number, result)
		fmt.Println("------------------")
	}

	fmt.Println("------------------")
	fmt.Println(result)

}
