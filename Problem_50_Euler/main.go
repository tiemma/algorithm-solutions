package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

// https://projecteuler.net/problem=49

func isPrime(n int) bool {
	sqrt := int(math.Sqrt(float64(n))) + 1
	for i := 2; i < sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func rotateDigit(n int) int {
	remainder := n - (10 * (n / 10))
	power := float64(len(strconv.Itoa(n))) - 1

	return (remainder * int(math.Pow(10, power))) + (n / 10)
}

func getCircularPrimes(n int) []int {
	newDigit := rotateDigit(n)
	values := []int{}
	count := 0

	for count != int(math.Log10(float64(n))+1) {
		newDigit = rotateDigit(newDigit)
		if newDigit < 1000 {
			return values
		}
		values = append(values, newDigit)
		count += 1
	}

	return values
}

func getMinimumValue(v []int) string {
	min := math.MaxInt32
	for _, v := range v {
		if v < min {
			min = v
		}
	}

	key := []rune(strconv.Itoa(min))
	sort.Slice(key, func(i int, j int) bool { return key[i] < key[j] })

	return string(key)
}

func dedup(v []int) []int {
	state := make(map[int]bool)
	var newV []int
	for _, k := range v {
		if _, ok := state[k]; ok {
			continue
		}
		state[k] = true
		newV = append(newV, k)
	}

	sort.Ints(newV)
	return newV
}

func main() {
	values := map[string][]int{}
	for i := 1000; i < 10000; i++ {
		if !isPrime(i) {
			continue
		}

		results := getCircularPrimes(i)
		sortingKey := getMinimumValue(results)

		if _, ok := values[sortingKey]; !ok {
			values[sortingKey] = []int{}
		}

		for _, value := range results {
			if !isPrime(value) {
				continue
			}
			values[sortingKey] = append(values[sortingKey], value)
		}
		values[sortingKey] = dedup(values[sortingKey])
	}

	for _, v := range values {
		diff := map[int][]int{}
		if len(v) < 3 {
			continue
		}
		for i, value := range v {
			for _, j := range v[i+1:] {
				key := j - value
				if _, ok := diff[key]; !ok {
					diff[key] = []int{}
				}
				diff[key] = append(diff[key], []int{j, value}...)
				diff[key] = dedup(diff[key])
			}
		}
		for k, v := range diff {
			sort.Ints(v)
			if len(v) == 3 {
				fmt.Println(k, v)
			}
		}
	}
}
