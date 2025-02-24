package main

// https://projecteuler.net/problem=18

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DFS(arr [][]int) int {
	for i := len(arr[len(arr)-1]) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if arr[i+1][j] > arr[i+1][j+1] {
				arr[i][j] += arr[i+1][j]
			} else {
				arr[i][j] += arr[i+1][j+1]
			}
		}
	}
	return arr[0][0]
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
	var triangleArr [][]int
	for buffer.Scan() {
		var numArr []int
		stringArr = strings.Split(strings.Trim(buffer.Text(), " "), ",")
		for _, el := range stringArr {
			val, _ := strconv.Atoi(el)
			numArr = append(numArr, val)
		}
		triangleArr = append(triangleArr, numArr)
	}
	fmt.Println(len(triangleArr[0]))
	result := DFS(triangleArr)
	fmt.Println(result)
}
