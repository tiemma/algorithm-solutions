package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://projecteuler.net/problem=59

func xor(a, b byte) byte {
	return a ^ b
}

func decrypt(text []byte, key []byte) []byte {
	var result []byte
	for i := 0; i < len(text); i++ {
		result = append(result, xor(text[i], key[i%len(key)]))
	}

	return result
}

func generateCombinations(numbers []string) [][]string {
	var result [][]string
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			for k := 0; k < len(numbers); k++ {
				result = append(result, []string{numbers[i], numbers[j], numbers[k]})
			}
		}
	}

	return result
}

func main() {
	fileName := flag.String("file", "0059_cipher.txt", "File to read integers from")
	flag.Parse()

	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	content, err := os.ReadFile(*fileName)
	if err != nil {
		panic(err)
	}

	var numbers []byte
	for _, line := range strings.Split(string(content), ",") {
		number, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, byte(number))
	}

	text, maxCount := "", 0
	for _, combination := range generateCombinations(letters) {
		var key []byte
		for _, c := range combination {
			key = append(key, []byte(c)[0])
		}

		result := decrypt(numbers, key)
		count := len(strings.Split(string(result), " "))

		if count > maxCount {
			text = string(result)
			maxCount = count
			fmt.Println(string(result), string(key), count)
		}
	}

	fmt.Println("------------")
	fmt.Println(text, maxCount)

	asciiSum := 0
	for _, c := range text {
		asciiSum += int(c)
	}

	fmt.Println(asciiSum)
}
