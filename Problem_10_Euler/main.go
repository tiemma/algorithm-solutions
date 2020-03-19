package main


// https://projecteuler.net/problem=10

import (
    "flag"
    "fmt"
    "math"
)


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

func sumOfPrimeToNBF(n int) int{
    i := 2
    sum := 0
    for i < n {
        if isPrime(i){
            sum += i
        }
        i++
    }    
    return sum
}

func main(){
    maxNumPtr := flag.Int("number", 10, "Integer for the max range")
    flag.Parse()
    fmt.Println("Prime sum to n is ", sumOfPrimeToNBF(*maxNumPtr))
}


