package main

// https://projecteuler.net/problem=38

import (
	"fmt"
	"math"
)

// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func getPermutations(arr []int)[][]int{
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int){
		if n == 1{
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++{
				helper(arr, n - 1)
				if n % 2 == 1{
					tmp := arr[i]
					arr[i] = arr[n - 1]
					arr[n - 1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n - 1]
					arr[n - 1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func getValue(digits []int)  int{
	value := 0
	for i := 0; i < len(digits); i++ {
		value += int(math.Pow(10, float64(i))) * digits[len(digits)-i-1]
	}

	return value
}

func getValueAndRemainder(x, y int) (int, int) {
	return x/y, x%y
}

func isPandigitalMultiple(digits []int) bool{
	numbers := []int{}
	for i := 1; i < 5; i++ {
		lastMultiple := 1
		divisor := getValue(digits[:i])
		numbers = append(numbers, divisor)
		newDigits := digits[i:]
		idx := i
		for len(newDigits) > 0 {
			if idx > len(newDigits) {
				break
			}
			multiple := getValue(newDigits[:idx])
			value, remainder := getValueAndRemainder(multiple, divisor)
			if value > lastMultiple + 1 {
				numbers = []int{}
				break
			}
			if !(value == lastMultiple + 1 && remainder == 0) {
				idx += 1
				continue
			}
			newDigits = newDigits[idx:]
			lastMultiple += 1
			numbers = append(numbers, multiple)

			if len(newDigits) == 0 {
				return true
			}
		}
	}

	return false
}


func main() {
	digits := []int{}
	for i := 1; i < 10; i++ {
		digits = append(digits, i)
	}

	value := math.MinInt64
	for _, numbers := range getPermutations(digits) {
		if isPandigitalMultiple(numbers) {
			newValue := getValue(numbers)
			fmt.Println(newValue)
			if value <  newValue{
				value = newValue
			}
		}
	}

	fmt.Println("--------------")
	fmt.Println(value)

}

