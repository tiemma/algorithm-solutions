
package main
// https://projecteuler.net/problem=15
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
    grid := flag.Int64("grid", 20, "Window of integers to  return")
    flag.Parse() 
 
    ud := factorial(big.NewInt(*grid * 2)) 
    ld :=  factorial(big.NewInt(*grid))
    n := big.NewInt(1)
    n.Div(ud, ld)
    n.Div(n, ld) 
    fmt.Println(n) 
}
