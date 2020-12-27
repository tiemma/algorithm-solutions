package main

// https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/529/week-2/3291/

import (
	"fmt"
	"strings"
)

func formatString(S string) string {
	s := []string{}
	for _, v := range S {
		n := len(s) - 1
		if string(v) == "#" {
			if n >= 0 {
				s = s[:n]
			}
		} else {
			s = append(s, string(v))
		}
	}
	return strings.Join(s, "")
}

func backspaceCompare(S string, T string) bool {
	return formatString(S) == formatString(T)
}
