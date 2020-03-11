package main


import (
    "fmt"
    "flag"
    "math"
)


// https://projecteuler.net/problem=5

func isPrime(n int) bool{
    i := 2
    for i <= int(math.Sqrt(float64(n))) {
        if n % i == 0 {
            return false
        }
        i++
    }
    return true
}


func smallestPositiveDivisibleNumber(nMax int) int {
   result := 1
    n := 2
    nMax--
    for n <= nMax {
        if isPrime(n) {
            fmt.Println("Is Prime ", n)
            invRoot := math.Pow(float64(n), -1)
            invPower := math.Floor(math.Pow(float64(nMax), invRoot))
            maxPower := math.Pow(float64(n), invPower) 
            fmt.Println("Largest multiple ", maxPower)
            result = result * int(maxPower)
        }
        n++
    }
    return result
}


func main(){
    maxNumPtr := flag.Int("number", 20, "Integer for the max range")
    flag.Parse()
    fmt.Println("Smallest divisible product ", smallestPositiveDivisibleNumber(*maxNumPtr))
}
