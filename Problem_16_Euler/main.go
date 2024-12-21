package main


// https://projecteuler.net/problem=16

import (
    "flag"
    "fmt"
    "math/big"
)

func main() {
    val := flag.Int64("val", 2, "Number to raise the power to")
    exp := flag.Int64("exp", 1_000, "Power of the value")
    flag.Parse()
    
    result := big.NewInt(0) 
    n := big.NewInt(1)
    n.Exp(big.NewInt(*val), big.NewInt(*exp), nil)
    for n.Cmp(big.NewInt(0)) != 0{
        temp := big.NewInt(0)
        n.DivMod(n, big.NewInt(10), temp)
        result.Add(result, temp)
    }
    fmt.Println("Sum is ", result)
}


