package Three_Number_Sum


import (
	"math"
	"sort"
)


// https://www.algoexpert.io/questions/Three%20Number%20Sum

func SumPow10(num int, num2 int, num3 int) int64 {
	return int64(math.Pow10(num) + math.Pow10(num2) + math.Pow10(num3))
}

func Find(slice []int64, val int64) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return  false
}

//Brute force O(mn3)
// m - Number of unique triplet combinations
// n - Number of elements in each array, equal sized
func ThreeNumberSum(array []int, target int) [][]int {
	sort.Ints(array)
	arr := [][]int{}
	var sums []int64
	for _, num := range array {
		for _, num2 := range array{
			for _, num3 := range array{
				sum := num + num2 + num3
				sumpow10 := SumPow10(num, num2, num3)
				if sum == target && !Find(sums, sumpow10) && num != num2 && num != num3 && num2 != num3{
					sums = append(sums,  sumpow10)
					arr = append(arr, []int{num, num2, num3})
				}
			}
		}
	}
	return arr
}


//Optimised version O(nlogn)
func ThreeNumberSum(array []int, target int) [][]int {
	sort.Ints(array)
	triplets := [][]int{}
	for idx, num := range array {
		left, right := idx+1, len(array) - 1;
		for left < right {
			currentSum := num + array[left] + array[right]
			if currentSum  == target{
				triplets  = append(triplets, []int{num, array[left], array[right]})
				left += 1
				right -= 1
			} else if currentSum > target{
				right -= 1
			} else if currentSum < target{
				left += 1
			}
		}
	}
	return triplets
}