package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
)

// https://projecteuler.net/problem=68

// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func getPermutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func getCombinations(arr, rest []int, n int, result [][]int) [][]int {
	if n == 0 {
		return [][]int{arr}
	} else {
		for i, k := range rest {
			result = append(result, getCombinations(append([]int{k}, arr...), rest[i+1:], n-1, [][]int{})...)
		}
	}

	return result
}

type entry struct {
	expectedSum int
	digits      [][]int
	usedDigits  map[int]struct{}
	result      [][]int
	finalValue  int
}

func filter(arr [][]int, filterMap map[int]struct{}, expectedSum int) [][]int {
	var newArr [][]int
	for i := 0; i < len(arr); i++ {
		_, ok1 := filterMap[arr[i][0]]
		_, ok2 := filterMap[arr[i][1]]
		if !ok1 && !ok2 && sum(arr[i]) == expectedSum {
			newArr = append(newArr, arr[i])
		}
	}

	return newArr
}

func sum(a []int) int {
	result := 0
	for i := 0; i < len(a); i++ {
		result += a[i]
	}
	return result
}

func backtrack(n int, digitMap map[int][][]int) []entry {
	var queue []entry
	firstEntry := entry{
		digits:      digitMap[n],
		usedDigits:  make(map[int]struct{}),
		result:      make([][]int, 0),
		expectedSum: 0,
	}
	queue = append(queue, firstEntry)

	var result []entry
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if len(cur.usedDigits) == len(digitMap) {
			if len(cur.digits) > 0 {
				panic(cur.digits)
			}
			resultLength := len(digitMap) / 2
			if cur.result[resultLength-1][1] == cur.result[resultLength-2][2] && cur.result[0][1] == cur.result[resultLength-1][2] {
				var nums []int
				for _, num := range cur.result {
					nums = append(nums, num...)
				}
				cur.finalValue = getValue(nums)
				result = append(result, cur)
			}
		}

		for _, d := range cur.digits {
			if len(cur.result) > 0 && getValue(cur.result[0]) > getValue(d) {
				continue
			}

			newUsedDigits := map[int]struct{}{}
			for k, v := range cur.usedDigits {
				newUsedDigits[k] = v
			}
			newUsedDigits[d[0]] = struct{}{}
			newUsedDigits[d[1]] = struct{}{}

			newResult := make([][]int, len(cur.result))
			copy(newResult, cur.result)
			newResult = append(newResult, d)

			newSum := sum(d)
			newDigits := filter(digitMap[d[2]], newUsedDigits, newSum)
			newEntry := entry{
				digits:      newDigits,
				usedDigits:  newUsedDigits,
				result:      newResult,
				expectedSum: newSum,
			}

			queue = append(queue, newEntry)
		}
	}

	return result
}

func getValue(digits []int) int {
	value := 0
	power := 0
	for i := 0; i < len(digits); i++ {
		number := float64(digits[len(digits)-i-1])
		value += int(math.Pow(10, float64(power)) * number)
		power += int(math.Floor(math.Log10(number)) + 1)
	}

	return value
}

func main() {
	nGonSize := flag.Int("d", 5, "N-gon ring to generate")
	expectedDigits := flag.Int("num", 16, "Expected number of digits to find")
	flag.Parse()

	nGonDigits := 2 * *nGonSize

	digits := []int{}
	for i := 1; i <= nGonDigits; i++ {
		digits = append(digits, i)
	}

	middleDigitMap := map[int][][]int{}
	for _, k := range getCombinations([]int{}, digits, 3, [][]int{}) {
		for _, d := range getPermutations(k) {
			middleDigitMap[d[1]] = append(middleDigitMap[d[1]], d)
		}
	}

	var result []entry
	for i := 1; i <= nGonDigits; i++ {
		result = append(result, backtrack(i, middleDigitMap)...)
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].expectedSum == result[j].expectedSum {
			return result[i].finalValue < result[j].finalValue
		}

		return result[i].expectedSum < result[j].expectedSum
	})

	maxNum := math.MinInt64
	for _, d := range result {
		numDigits := int(math.Log10(float64(d.finalValue))) + 1
		if d.finalValue > maxNum && numDigits == *expectedDigits {
			maxNum = d.finalValue
		}
		fmt.Println(d.expectedSum, d.result, d.finalValue, int(math.Log10(float64(d.finalValue)))+1)
	}
	fmt.Println("-------------")
	fmt.Println(maxNum)
}
