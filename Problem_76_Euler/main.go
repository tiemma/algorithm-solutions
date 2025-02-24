package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
	"strconv"
)

// https://projecteuler.net/problem=76

func getPentagonalNumber(n int) int {
	return n * (3*n - 1) / 2
}

func computePartitions(goal int) []int {
	partitions := []int{1}
	for n := 1; n <= goal+1; n++ {
		partitions = append(partitions, 0)
		for k := 1; k <= n; k++ {
			coefficient := int(math.Pow(-1, float64(k+1)))
			for _, t := range []int{getPentagonalNumber(k), getPentagonalNumber(-k)} {
				if (n - t) >= 0 {
					partitions[n] += coefficient * partitions[n-t]
				}
			}
		}
	}
	return partitions
}

func main() {
	window := flag.Int("window", 100, "The window size")
	flag.Parse()

	partitions := computePartitions(*window)
	//https://en.wikipedia.org/wiki/Partition_function_(number_theory)#Recurrence_relations
	// We subtract 1 from the answer as it also includes the single value of window which is exempt in project euler
	fmt.Println(partitions)
	fmt.Println("************")
	fmt.Println(partitions[*window] - 1)

}

func getKey(n []int) string {
	result := ""
	for _, i := range n {
		result += "-" + strconv.Itoa(i)
	}

	return result
}

func bruteForce(n int) int {
	var starting []int
	for len(starting) < n {
		starting = append(starting, 1)
	}
	found := map[string]bool{}
	counters := map[int]int{}

	queue := make([][]int, 0)
	queue = append(queue, starting)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		fmt.Println(cur)
		if len(cur) == 1 || cur[0] != 1 {
			continue
		}
		counters[len(cur)] += 1

		for j := 1; j < len(cur); j++ {
			newCur := make([]int, len(cur))
			copy(newCur, cur)
			newCur[j] += 1
			anotherCur := newCur[1:]
			sort.Ints(anotherCur)

			key := getKey(anotherCur)
			if _, ok := found[key]; ok {
				continue
			}
			found[key] = true

			queue = append(queue, anotherCur)
		}
	}

	fmt.Println(n, counters, len(found))
	return len(found)
}
