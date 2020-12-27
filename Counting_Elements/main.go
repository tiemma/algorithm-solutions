package main

// https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/528/week-1/3289/

import (
	"fmt"
	"sort"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func countElements(arr []int) int {
	setElem := map[int]int{}
	count := 0
	temp := 0
	for _, val := range arr {
		if _, ok := setElem[val]; !ok {
			setElem[val] = 0
		}
		setElem[val] += 1
	}

	keys := []int{}
	for k, _ := range setElem {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, val := range keys {
		if _, ok := setElem[val+1]; ok {
			temp += setElem[val]
		} else {
			count += temp
			temp = 0
		}
	}
	return count
}
