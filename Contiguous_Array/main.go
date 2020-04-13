package main

// https://leetcode.com/problems/contiguous-array

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxLength(nums []int) int {
	max_len, sum := 0, 0
	count := map[int]int{}
	count[0] = -1
	for idx, val := range nums {
		adder := 1
		if val == 0 {
			adder = -1
		}
		sum += adder
		if elem, ok := count[sum]; ok {
			max_len = max(max_len, idx-elem)
		} else {
			count[sum] = idx
		}
	}

	return max_len
}
