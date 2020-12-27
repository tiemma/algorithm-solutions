package main

import (
    "flag"
    "fmt"
    "math"
)

// https://projecteuler.net/problem=24

func swap(arr []int, i int, j int) {
    temp := arr[i]
    arr[i] = arr[j]
    arr[j] = temp
}

func Reverse(s []int) []int {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    } 
    return s 
}


func Equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

func factorial(x int) int {
    n := x
    if x == 0 {
        return 1
    }
    return n * factorial(x-1)
}

func main(){
    window := flag.Int("window", 2, "Length of the number to generate to")
    index := flag.Int("index", 5, "Nth lexicographic permutation to look out for")
    flag.Parse()

    arr := []int{}
    i := 0
    for i <= *window {
        arr = append(arr, i)
        i++
    }
   
    fmt.Println(arr) 
    i = 1
    fmt.Println(factorial(*window+1))
    for i < factorial(*window + 1) {
        // Find the largest X
        largestI := -1
        for idx, val := range arr {
            if idx < len(arr) - 1 && val < arr[idx + 1] {
                largestI = idx
            }
        }

    
        largestJ := -1
        for idx, val :=  range arr {
            if arr[largestI] < val {
                largestJ = idx
            }
        }
    
        swap(arr, largestI, largestJ)

        sliceArr := Reverse(arr[largestI+1:])
    
        arr = append(arr[:largestI+1], sliceArr...)
        // fmt.Println(arr)
        if i == *index - 1 {
            number := 0
            for idx, val := range arr {
                number += int(math.Pow(10, float64(len(arr) - idx - 1))) * val
            }
            fmt.Println(number)
            break
        }
        i += 1 
    }
}
