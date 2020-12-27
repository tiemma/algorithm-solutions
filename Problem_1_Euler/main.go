package main


// https://projecteuler.net/problem=1

import "fmt"


func calcArithmeticSum(a int, n int) int {
	if n == 0 {
		return 0
	}
	ans := n * (a * (n+1)) / 2
	fmt.Printf("Answer with a:%d and n:%d is %d\n", a, n , ans)
	return ans
}

// O(1) space and O(1) time
func MultiplesOf3And5(num int) int {
	divBy3 := (num - 1) / 3
	divBy5 := (num - 1) / 5
	divBy15 := (num - 1) / 15
	return calcArithmeticSum(3, divBy3) + calcArithmeticSum(5, divBy5) - calcArithmeticSum(15, divBy15)
}


func main() {
	num := 1000
	ans := MultiplesOf3And5(num)
	fmt.Println(ans)
}