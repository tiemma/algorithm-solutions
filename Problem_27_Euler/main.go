package main

import (
    "fmt"
    "flag"
    "math"
)

// https://projecteuler.net/problem=27

func isPrime(n float64, a float64, b float64) bool{
    val := math.Pow(n, 2) + (a * n) + b 
    var i float64 = 2
    if val <= 1 {
        return false
    }
    for i <= math.Sqrt(val) {
        if int(val) % int(i) == 0 {
            return false
        }
        i += 1
    }
    return true
}


func main() {
    vala := flag.Int("a", 1_000, "Range of values for a")
    valb := flag.Int("b", 1_000, "Range of values for b")
    flag.Parse()

    a := - float64(*vala)
    var max float64 = 1
    var result float64 = 1
    fmt.Println(result)
    for a <= float64(*vala) {        
        b := - float64(*valb)
        for b <= float64(*valb) {
            var n float64 = 0
            for isPrime(n, a, b) {
                n += 1
            }
            /*
            if n > 1 {
                fmt.Println(n - 2, a, b)
            }*/
            if n - 2 > max {
                max = n - 2
                result = a * b
            }      
            b += 1
        }
        a += 1
    }
    fmt.Println(result) 
}
