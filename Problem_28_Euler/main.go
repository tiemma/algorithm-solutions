package main

import (
  "flag"
  "fmt"
)

// https://projecteuler.net/problem=28


func diagonalDifferenceFactor(num int) int {
  return num - 1
}

func topRight(num int) int {
  return num * num
}

func topLeft(num int) int {
  return topRight(num) - diagonalDifferenceFactor(num) 
}

func bottomLeft(num int) int {
  return topRight(num) - (4 * (diagonalDifferenceFactor(num) / 2))
}

func bottomRight(num int) int { 
  return topRight(num) - (3 * diagonalDifferenceFactor(num))
}

func main() {
  grid := flag.Int("grid", 1_001, "Spiral width along the diagonal")
  flag.Parse()

  sum := 1
  i := 3
  for i <= *grid{
    sum += topRight(i) + topLeft(i) + bottomLeft(i) + bottomRight(i)
    i += 2
  }
  fmt.Println(sum)
}
