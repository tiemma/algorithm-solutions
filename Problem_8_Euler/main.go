package main

import (
    "math"
    "fmt"
    "flag"
    "strconv"
)


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
    elems := flag.String("elems", "20", "Integer element to run the product over")
    window := flag.Int("window", 20, "Integer for the max product window")
    flag.Parse()
    fmt.Println(loopAcrossValueSets(*elems, *window))

}

