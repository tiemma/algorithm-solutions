package main

import (
	"fmt"
	"math"
)

func BinetFormula(n float64) int64 {
	sqrt5 := math.Sqrt(5)
	phi := (1 + sqrt5) / 2
	ans := math.Round(math.Pow(phi, n) / sqrt5)
	fmt.Printf("Value at n: %f in fibonnaci sequence is: %f\n", n, ans)
	return int64(ans)
}

func LastDigit(n float64) int64 {
	var seqMax int64 = 60
	var count int64 = 0
	var val int64 = 0

	fibSequence := make([]int64, seqMax)
	for count < seqMax {
		if count == 0{
			fibSequence[count] = 0
		}
		fibSequence[count] = (fibSequence[count] + BinetFormula(float64(count))) % 10
		count += 1
	}

	count = 0
	for count <= int64(n) % 60{
		val += fibSequence[count]
		count += 1
	}

	return val % 10
}


func main(){
	fmt.Println(LastDigit(832564823476))
}