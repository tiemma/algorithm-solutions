package main

// https://projecteuler.net/problem=81

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func DFS(arr [][]float64) float64 {
	var newArr [][]float64
	size := len(arr)
	for i := 0; i < len(arr); i++ {
		var entry []float64
		for j := 0; j < len(arr[i]); j++ {
			entry = append(entry, 0)
		}
		newArr = append(newArr, entry)
	}
	newArr[0][0] = arr[0][0]

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Println("-----------")

	for i := 0; i <= len(arr)-1; i++ {
		for j := 0; j <= len(arr[i])-1; j++ {
			current := newArr[i][j]

			if i+1 < size {
				newY := current + arr[i+1][j]
				if newArr[i+1][j] > 0 {
					newY = math.Min(newArr[i+1][j], newY)
				}
				newArr[i+1][j] = newY
			}

			if j+1 < size {
				newX := current + arr[i][j+1]
				if newArr[i][j+1] > 0 {
					newX = math.Min(newArr[i][j+1], newX)
				}
				newArr[i][j+1] = newX
			}
		}
	}

	for i := 0; i < size; i++ {
		fmt.Println(newArr[i])
	}
	fmt.Println("-----------")

	return newArr[size-1][size-1]
}

func main() {
	fileName := flag.String("file", "0081_matrix.txt", "File to read integers from")
	flag.Parse()

	input, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	buffer := bufio.NewScanner(input)
	var stringArr []string
	var triangleArr [][]float64
	for buffer.Scan() {
		var numArr []float64
		stringArr = strings.Split(strings.Trim(buffer.Text(), " "), ",")
		for _, el := range stringArr {
			val, _ := strconv.ParseFloat(el, 64)
			numArr = append(numArr, val)
		}
		triangleArr = append(triangleArr, numArr)
	}
	fmt.Println(len(triangleArr[0]))
	result := DFS(triangleArr)
	fmt.Println(result)
}
