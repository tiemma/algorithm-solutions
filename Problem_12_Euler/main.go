package main


// https://projecteuler.net/problem=12

import (
    "flag"
    "fmt"
    "math"
)

type Maps map[int]map[int]int

// The triangle number sum is a simple AP sum 
func getTriangleNumber(n int) int { 
    return  n * (n + 1) / 2
}


// Every single natural number is a mutiple of a prime
// So we can decompose that product with a simple prime factor check
// And reduce the powers till we get another prime
func countMultiples(n int, divisorMap Maps) map[int]int {
    i := 2
    factorMap := map[int]int{}
    temp := n

    for i <= int(math.Sqrt(float64(n))) {
        for n % i == 0 && n > 1 {
            n = n /i  
            
            if val, ok := factorMap[i]; !ok{
                factorMap[i] = 1
            } else {
                factorMap[i] = val + 1   
            }

            if _, ok := divisorMap[n]; ok {
                fmt.Println("Found value for number ", n, " in divisor map for ", temp)
                for key, val := range divisorMap[n] {
                    if _, ok := factorMap[key]; !ok {
                        factorMap[key] = val
                    } else {
                        factorMap[key] += val
                    }
                }
                return factorMap
            }
 
        }
        i++
    }

    if n > 1 {
        factorMap[n] = 1
    
    }

    return factorMap
}


func highestDivisibleTriangularNumber(maxDivisor int) int{
    divisorMap := Maps{}
    n := 2
    for n > 0{
        temp := getTriangleNumber(n)
        divisorMap[temp] = countMultiples(temp, divisorMap) 
        // The number of factor compositions is a product of the order of each prime
        product := 1
        for _,v := range divisorMap[temp]{    
            product *= (v + 1)
        }
        if product >= maxDivisor{
            return temp
        }
        n += 1
    }
    return 1
}


func main(){
    //divisorMap := Maps{15: {3: 1, 5: 1}} 
    divisor := flag.Int("divisor", 2, "Flag to indicate the number of divisors to look out for")
    flag.Parse()
    //fmt.Println("Number of multiples to ", *divisor, " is ", countMultiples(*divisor, divisorMap))
    fmt.Println("Number of multiples to ", *divisor, " is ", highestDivisibleTriangularNumber(*divisor))
}


