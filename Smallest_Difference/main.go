package Smallest_Difference

import (
"sort"
"math"
)

// https://www.algoexpert.io/questions/Smallest%20Difference

// Brute Force O(mn)
// n - len(array1)
// m - len(array2)
func SmallestDifference(array1 []int, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)
	arr, diff := []int{}, math.Inf(1)

	for _, num := range array1 {
		for _, num2 := range array2 {
			if math.Abs(float64(num - num2)) < diff {
				diff = math.Abs(float64(num - num2))
				arr = []int{num, num2}
			}
		}
	}

	return arr
}


// Optimised version O(nlogn + mlogm)
// n - len(array1)
// m - len(array2)
func SmallestDifference(array1 []int, array2 []int) []int {
	// Write your code here.
	sort.Ints(array1)
	sort.Ints(array2)
	arr, diff := []int{}, math.MaxInt64
	idx1, idx2, current := 0, 0, 0
	for idx1 < len(array1) && idx2 < len(array2){
		first, second := array1[idx1], array2[idx2]
		if first > second {
			current = first - second
			idx2+=1
		} else if second > first {
			current = second - first
			idx1+=1
		} else {
			return []int{first, second}
		}
		if diff > current{
			diff = current
			arr = []int{first, second}
		}
	}

	return arr
}