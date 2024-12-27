package main

import (
	"flag"
	"fmt"
	"math"
	"sync"
)

// https://projecteuler.net/problem=60

func isPrime(n int) bool {
	i := 2
	if n == 0 || n == 1 {
		return false
	}
	for i <= int(math.Sqrt(float64(n))) {
		if n%i == 0 {
			return false
		}
		i++
	}

	return true
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

func isPrimePair(a, b int) bool {
	for _, number := range [][]int{{a, b}, {b, a}} {
		if !isPrime(getValue(number)) {
			return false
		}
	}

	return true
}

func sum(n []int) int {
	result := 0
	for _, n := range n {
		result += n
	}

	return result
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

func processResult(combChan, resultChan chan []int, primes []int, wg *sync.WaitGroup) {
	for comb := range combChan {
		for _, p := range primes {
			if comb[0] >= p {
				continue
			}
			isValid := true
			newList := append([]int{p}, comb...)
			for _, d := range getCombinations([]int{}, newList, 2, [][]int{}) {
				if !isPrimePair(d[0], d[1]) {
					isValid = false
					break
				}
			}
			if isValid {
				resultChan <- newList
			}
		}
	}
	wg.Done()
}

func main() {
	window := flag.Int("window", 4, "Number of prime factors to find replacements matching")
	flag.Parse()

	for primeWindow := 1000; primeWindow < math.MaxInt64; primeWindow += 1000 {
		fmt.Println("Prime window:", primeWindow)
		var primes []int
		for i := 1; i < primeWindow; i++ {
			if isPrime(i) {
				primes = append(primes, i)
			}
		}

		fmt.Println(len(primes), *window)

		combinations := getCombinations([]int{}, primes, 5, [][]int{})
		fmt.Println(len(combinations), *window)
		fmt.Println("-------------")

		//for i := 3; i < *window; i++ {
		//	var newCombinations [][]int
		//	for _, comb := range combinations {
		//		for _, p := range primes {
		//			if comb[0] >= p {
		//				continue
		//			}
		//			isValid := true
		//			newList := append([]int{p}, comb...)
		//			for _, d := range getCombinations([]int{}, newList, 2, [][]int{}) {
		//				if !isPrimePair(d[0], d[1]) {
		//					isValid = false
		//					break
		//				}
		//			}
		//			if isValid {
		//				newCombinations = append(newCombinations, newList)
		//				if i == *window {
		//					fmt.Println("-------------")
		//					fmt.Println(newList, sum(newList))
		//					return
		//				}
		//			}
		//		}
		//	}
		//}

		for i := 3; i < *window+1; i++ {
			var newCombinations [][]int
			combChan := make(chan []int)
			resultChan := make(chan []int)
			wg := &sync.WaitGroup{}
			for k := 0; k < 15; k++ {
				wg.Add(1)
				go processResult(combChan, resultChan, primes, wg)
			}
			go func() {
				for _, k := range combinations {
					combChan <- k
				}
				close(combChan)
				wg.Wait()
				close(resultChan)
			}()

			for newList := range resultChan {
				newCombinations = append(newCombinations, newList)
				if i == *window {
					fmt.Println("-------------")
					fmt.Println(newList, sum(newList))
					return
				}
			}
			combinations = newCombinations
			fmt.Println(len(newCombinations), i)
		}
	}
}
