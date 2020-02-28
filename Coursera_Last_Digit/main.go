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
	fibSequence := make([]int64, seqMax)
	var count int64 = 1
	for count <= seqMax {
		fibSequence = append(fibSequence, BinetFormula(float64(count))  % 10)
		count += 1
	}
	return fibSequence[int(64) % 60]
}


func main(){
	fmt.Println(LastDigit(832564823476))
}