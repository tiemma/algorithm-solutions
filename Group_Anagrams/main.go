package main

// https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/528/week-1/3288/

import (
	"sort"
	"strings"
)

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func groupAnagrams(strs []string) [][]string {
	content := map[string][]string{}
	output := [][]string{}
	for _, str := range strs {
		idx := sortString(str)
		content[idx] = append(content[idx], str)
	}

	for _, v := range content {
		output = append(output, v)
	}
	return output
}
