package main

// https://leetcode.com/problems/longest-palindromic-substring/

import (
	"strings"
)

func oddifyString(s string) []string {
	val := []string{}
	for _, s := range s {
		val = append(val, "#")
		val = append(val, string(s))
	}
	val = append(val, "#")
	return val
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func longestPalindrome(sb string) string {
	s := oddifyString(sb)
	N := len(s)
	res := []string{}
	if N == 0 {
		return ""
	}
	p := make([]int, N)
	c := 0
	r := 0
	max := 0
	for i := 0; i < N; i++ {
		mirror := (2 * c) - i
		if i < r {
			p[i] = min(r-i, p[mirror])
		}

		a := i + (1 + p[i])
		b := i - (1 + p[i])

		for a < N && b >= 0 && s[a] == s[b] {
			p[i]++
			a++
			b--
		}

		if i+p[i] > r {
			c = i
			r = i + p[i]

			if p[i] > max {
				max = p[i]
				res = s[i-p[i] : r+1]
			}
		}
	}
	return strings.ReplaceAll(strings.Join(res, ""), "#", "")
}
