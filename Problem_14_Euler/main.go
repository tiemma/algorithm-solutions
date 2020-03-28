package main


import (
    "flag"
    "fmt"
    "math"
)

type Map map[int]int

func getCollatzSequenceCount(n int, countMap Map) int{
    count := 1
    for n > 1 {
        if val, ok := countMap[n]; ok {
            return count + val
        }
        if n%2 == 1{
            n = 3*n + 1
        } else {
            n /= 2
        }
        count += 1
    }
    return count
}

func getTermsCount(maxCount int) int{
    count := math.MinInt64
    n := 1
    result := 0
    countMap := Map{}
    for n < maxCount{
        temp := getCollatzSequenceCount(n, countMap)  
        countMap[n] = temp
        if count < temp {
            count  = temp
            result = n
        } 
        n += 1     
    }
    return result
}


func main(){
    maxTerms := flag.Int("max", 10, "Max number of terms obtained from the collatz sequence")
    flag.Parse()

   fmt.Println("Variable with max iterations is ", getTermsCount(*maxTerms))
}
