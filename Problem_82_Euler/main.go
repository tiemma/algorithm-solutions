package main

// https://projecteuler.net/problem=82

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
	var queue [][3]float64
	var newArr [][]float64
	size := float64(len(arr))
	for i := 0; i < len(arr); i++ {
		var entry []float64
		for j := 0; j < len(arr[i]); j++ {
			entry = append(entry, 0)
		}
		newArr = append(newArr, entry)
	}

	for i := 0; i < len(arr); i++ {
		queue = append(queue, [3]float64{float64(i), float64(0), arr[i][0]})
	}

	globalMin := math.MaxFloat64
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		x, y, val := cur[0], cur[1], cur[2]
		if x >= size || y >= size || x < 0 {
			continue
		}

		if newArr[int(x)][int(y)] == 0 || newArr[int(x)][int(y)] > val {
			newArr[int(x)][int(y)] = val
		} else {
			continue
		}

		if y == size-1 {
			globalMin = math.Min(globalMin, val)
		}

		if x-1 >= 0 {
			queue = append(queue, [3]float64{x - 1, y, val + arr[int(x-1)][int(y)]})
		}

		if x+1 < size {
			queue = append(queue, [3]float64{x + 1, y, val + arr[int(x+1)][int(y)]})
		}

		if y+1 < size {
			queue = append(queue, [3]float64{x, y + 1, val + arr[int(x)][int(y+1)]})
		}
	}

	return globalMin
}

func main() {
	fileName := flag.String("file", "0082_matrix.txt", "File to read integers from")
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
