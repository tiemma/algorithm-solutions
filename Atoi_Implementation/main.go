package main

// https://leetcode.com/problems/string-to-integer-atoi/

import (
	"math"
	"strconv"
	"strings"
)

func myAtoi(str string) int {
	str = strings.TrimLeft(str, " ")
	str = strings.Split(str, " ")[0]
	str = strings.Split(str, ".")[0]
	elem := []string{}
	for idx, val := range str {
		if idx == 0 && strings.ContainsAny(string(val), "-+") {
			elem = append(elem, string(val))
			continue
		}
		if _, err := strconv.Atoi(string(val)); err == nil {
			elem = append(elem, string(val))
		} else {
			break
		}
	}
	str = strings.Join(elem, "")
	val, err := strconv.ParseInt(str, 10, 64)
	if val < 0 && val < math.MinInt32 {
		return math.MinInt32
	} else if val > math.MaxInt32 {
		return math.MaxInt32
	}
	if err == nil {
		return int(val)
	}
	return 0
}
