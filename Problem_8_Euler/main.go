package main

import (
    "math"
    "fmt"
    "flag"
    "os"
    "strconv"
)

// https://projecteuler.net/problem=8


func getSliceProduct(elems []string) int{
    result := 1
    for _, i := range elems {
        temp, err := strconv.Atoi(i)
        if err != nil{
            fmt.Println("Element is not integer ", i, ", ignored during loop run")
        }
        result *= temp
    }
    return result
}


func loopAcrossValueSets(elemStr string, window int) int{
    i := 0
    elems := []string{}
    for i, _ := range elemStr {
        elems = append(elems, string(elemStr[i]))
    }
    result := math.MinInt64
    fmt.Println(elems)
    for i < (len(elems) - window) {
        temp := getSliceProduct(elems[i:i+window])
        if temp > result {
            result = temp
        } 
        i++
    }
    return result
} 

func main(){
    content, err := os.ReadFile("number.txt")
    if err != nil {
        panic(err)
    }
    elems := flag.String("elems", string(content), "Integer element to run the product over")
    window := flag.Int("window", 13, "Integer for the max product window")
    flag.Parse()
    fmt.Println(loopAcrossValueSets(*elems, *window))

}

