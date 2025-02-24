package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// https://projecteuler.net/problem=79

func main() {
	fileName := flag.String("file", "0079_keylog.txt", "The name of the keylog file")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	numbers := map[string]map[string]bool{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var arr []string
		for _, index := range line {
			if numbers[string(index)] == nil {
				numbers[string(index)] = map[string]bool{}
			}
			for _, d := range arr {
				numbers[string(index)][d] = true
			}
			arr = append(arr, string(index))
		}
	}

	var digits []string
	for char, _ := range numbers {
		digits = append(digits, char)
	}

	fmt.Println(digits)
	fmt.Println(numbers)
	sort.Slice(digits, func(i, j int) bool {
		a, b := digits[i], digits[j]

		return len(numbers[a]) < len(numbers[b])
	})
	fmt.Println("**********")
	fmt.Println(strings.Join(digits, ""))
}
