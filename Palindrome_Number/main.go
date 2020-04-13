package main

// https://leetcode.com/problems/palindrome-number/

import (
	"strconv"
	"strings"
)

func Reverse(s string) string {
	elem := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		elem = append(elem, string(s[i]))
	}
	return strings.Join(elem, "")
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	strNum := strconv.Itoa(x)
	if strNum == Reverse(strNum) {
		return true
	}
	return false
}
