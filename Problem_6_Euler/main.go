package main


import (
    "flag"
    "fmt"
)


// https://projecteuler.net/problem=6

func getLinearAPSum(n int) int{
    return (n * (n + 1)) * (n * (n + 1)) / 4
}

func getSquareAPSum(n int) int{
    return ( n * (n + 1) * ((2 *n) + 1) ) / 6
}

func main() {
    maxNumPtr := flag.Int("number", 100, "Integer for the max range")
    flag.Parse()
    fmt.Println(-getSquareAPSum(*maxNumPtr) + getLinearAPSum(*maxNumPtr))
}
