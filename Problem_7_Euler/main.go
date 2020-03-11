package main

import (
    "flag"    
    "fmt"
    "math"
)

// https://projecteuler.net/problem=7

// We can offset unncessary computation by setting the prime check startpoint
// using the prime count function
// https://mathworld.wolfram.com/PrimeNumberTheorem.html
// For simplicity, the table listed in the @findPrimeAtIndex function is used to approximate the occurrence
// Additional study on implementing the code to generate the table is for the viewers discretion to do so


// Not too efficient, overflows on n=17
// Implements Wilson's theorem on prime validation
// Implementation guide: https://codegolf.stackexchange.com/a/94385
func primeCountingFunction(n int64, k int64, p int64) int{
    count := 0
    for k <= n{
        // A prime is valid if p%k == 1
        // and not if p%k == 0
        if p % k == 1 {
            count ++
        }
        fmt.Println(k, p, count)

        p *= k * k
        k += 1
    }
    return count
}


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



func findPrimeAtIndex(n int) int{
    // Using prime counting function table values
    // Full table available here: https://mathworld.wolfram.com/PrimeCountingFunction.html
   // Each index here corresponds to a count up to 10 of the power of that index
   // 1 - 10, 2 -100, 3 - 1000 etc.
    primeCount := []int{4, 25, 168, 1229, 9592, 78498, 664579}
    availPrimeCount := 4
    nCountValue := 10

    for idx, num := range primeCount {
        if n < num{
            break
        }
        availPrimeCount = num
        nCountValue = int(math.Pow(10, float64(idx+1)))
    }
    
    fmt.Println("Prime count, Current Value: ", availPrimeCount, nCountValue)

    for n > availPrimeCount {
        nCountValue++
        if isPrime(nCountValue) {
            availPrimeCount++
        }
    }

    return nCountValue
}

func main(){
    maxNumPtr := flag.Int("number", 20, "Integer for the max range")
    flag.Parse()
    fmt.Println("Prime is ", findPrimeAtIndex(*maxNumPtr))
}


