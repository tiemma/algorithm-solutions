package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

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

func bruteforce(window int) int {
	n := 1
	minN := float64(math.MaxInt64)
	var primes []int

	for i := 2; i <= window; i++ {
		count := float64(i)
		if isPrime(i) {
			primes = append(primes, i)
			count -= 1
			continue
		} else {
			// https://en.wikipedia.org/wiki/Euler%27s_totient_function#Computing_Euler's_totient_function
			for _, p := range primes {
				if i%p == 0 {
					count *= 1 - (1 / float64(p))
				}
			}
		}

		a := strings.Split(strconv.Itoa(int(count)), "")
		b := strings.Split(strconv.Itoa(i), "")
		sort.Strings(a)
		sort.Strings(b)

		if strings.Join(a, "") == strings.Join(b, "") {
			fmt.Println(i, int(count))
			phiN := float64(i) / count
			if phiN < minN {
				minN = phiN
				n = i
			}
		}
	}

	fmt.Println("------------------")
	fmt.Println(n)

	return n
}

func main() {
	window := flag.Int("window", 10_000_000, "The window size")
	flag.Parse()

	// https://mathworld.wolfram.com/TotientFunction.html
	// The idea for the new solution is that a permutation close to n will happen when the prime is close to the count
	// We already know it will be close to prime as the phi(n) = n - 1
	// To avoid searching for those values, we assume that n is a product of two primes
	// therefore if n = p*q, then phi(n) = phi(pq) = (p-1)(q-1)
	// The entire body of this solution is to find the two primes that fit within n < N
	// and then use the value to get the solution that minimises the n/phi(n)
	// We cannot use a prime as we already know that n-1 compared to n will not permute but will be minimal for n/phi(n)

	sqrtN := int(math.Sqrt(float64(*window)))
	primeRange := int(math.Pow(10, math.Floor(math.Log10(float64(sqrtN))+2)))
	var primes []int
	minPhi := float64(math.MaxInt64)
	maxN := 0
	for i := sqrtN - primeRange; i < sqrtN+primeRange; i++ {
		if i < 0 {
			continue
		}
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	for i := 0; i < len(primes); i++ {
		for j := i + 1; j < len(primes); j++ {
			prod := primes[i] * primes[j]
			if prod > *window {
				break
			}

			phiN := (prod + 1) - (primes[i] + primes[j]) // Expanded version of (p-1)(q-1)
			a := strings.Split(strconv.Itoa(prod), "")
			b := strings.Split(strconv.Itoa(phiN), "")
			sort.Strings(a)
			sort.Strings(b)

			if strings.Join(a, "") == strings.Join(b, "") {
				ratio := float64(prod) / float64(phiN)

				fmt.Println(prod, phiN, ratio)
				if ratio < minPhi {
					minPhi = ratio
					maxN = prod
				}
			}
		}
	}

	fmt.Println("------------------")
	fmt.Println(maxN)
}
