package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// https://projecteuler.net/problem=33

func splitInt(n int) []string {
	return strings.Split(strconv.Itoa(n), "")
}

func hasPanFactors(n int) bool {
	for i := 1; i * i <= n; i++ {
		s := splitInt(n)

		if n % i == 0 {
			s = append(s, splitInt(i)...)
			s = append(s, splitInt(n / i)...)
		}

		sort.Strings(s)

		if strings.Join(s, "") == "123456789" {
			return true
		}
	}

	return false
}

func main() {
	sum := 0
	for i := 1; i < 10000; i++ {
		if hasPanFactors(i) {
			sum += i
		}
	}

	fmt.Println("=====")
	fmt.Println(sum)
}