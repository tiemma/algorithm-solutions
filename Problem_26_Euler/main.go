package main


import (
    "flag"
    "fmt"
)


// https://projecteuler.net/problem=26

func getRepetend(numer int, denom int) int {
    pos := -1
    base := 10
    digitMap := map[int]int{}
    for {
        
        if _, ok := digitMap[numer]; ok || numer == 0 {
            break
        }
        digitMap[numer] = pos
        z := base * numer / denom
        numer = (numer * base) - (z * denom)
        fmt.Println(numer, pos)
        pos -= 1
    }

    return digitMap[numer] - pos
}


func main(){
    window := flag.Int("window", 1_000, "Range to find the longest repetend in")
    flag.Parse()
   
    result, max := 0, 0
    for *window > 1  {
        *window --
        temp := getRepetend(1, *window)
        if temp > max {
            result = * window
            max = temp
        } 
        *window -- 
    }
    fmt.Println("Digit, Length: ",result, max) 
}

