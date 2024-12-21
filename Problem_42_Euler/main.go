package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

// https://projecteuler.net/problem=42

func isTriangleWord(c int) bool {
	a := 1.0
	b := 1.0
	firstRoot := (-b + math.Sqrt(math.Pow(b, 2)-(4*a*float64(c)))) / (2 * a)
	secondRoot := (-b - math.Sqrt(math.Pow(b, 2)-(4*a*float64(c)))) / (2 * a)

	return math.Floor(firstRoot) == firstRoot && math.Floor(secondRoot) == secondRoot
}

func convertWordToNumber(word string) int {
	result := 0

	for i := 1; i < len(word)-1; i++ {
		result += int(word[i]-'A') + 1
	}

	return result
}

func main() {
	content, err := os.ReadFile("words.txt")
	if err != nil {
		panic(err)
	}

	// In this case, we use prime factorization via the quadratic formula
	// ax^2 + bx + c = 0
	// n/2(n+1) = c
	// n^2 + n  = 2c
	// n^2 + n - 2c = 0
	// All that's done here is converting the number to an int and multiplying the value by -2
	// To check if it's a triangle word is just simple arithmetic to verify the roots are whole numbers
	lines := strings.Split(string(content), ",")
	count := 0
	for _, line := range lines {
		if isTriangleWord(-2 * convertWordToNumber(line)) {
			count += 1
		}
	}

	fmt.Println(count)
}
