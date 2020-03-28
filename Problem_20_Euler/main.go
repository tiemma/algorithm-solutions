package main

import (
    "flag"
    "fmt"
    "math/big"
)

func factorial(x *big.Int) *big.Int {
    n := big.NewInt(1)
    if x.Cmp(big.NewInt(0)) == 0 {
        return n
    }
    return n.Mul(x, factorial(n.Sub(x, n)))
}


func main() {
    grid := flag.Int64("grid", 10, "Window of integers to  return")
    flag.Parse() 
 
    ud := factorial(big.NewInt(*grid)) 
    result := big.NewInt(0)
    for ud.Cmp(big.NewInt(0)) != 0{
        temp := big.NewInt(0)
        ud.DivMod(ud, big.NewInt(10), temp)
        result.Add(temp, result)
    }

    fmt.Println(result) 
}
