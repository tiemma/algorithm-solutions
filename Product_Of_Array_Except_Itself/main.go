package main

// https://leetcode.com/problems/product-of-array-except-self/

func productExceptSelf(nums []int) []int {
	arr := make([]int, len(nums))
	p := 1
	for idx, el := range nums {
		arr[idx] = 1
		arr[idx] *= p
		p *= el
	}

	p = 1
	for idx := len(nums) - 1; idx >= 0; idx-- {
		arr[idx] *= p
		p *= nums[idx]
	}
	return arr
}
