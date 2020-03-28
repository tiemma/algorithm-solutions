package main


// https://projecteuler.net/problem=22

import (
    "bufio"
    "fmt"
    "flag"
    "os"
    "sort"
    "strings"
)




func main(){
    fileName := flag.String("file", "input.txt", "File to read names from")
    flag.Parse()

    input, err := os.Open(*fileName)
    if err != nil {
        panic(err)
    }
    defer input.Close()

    buffer := bufio.NewScanner(input)
    stringArr := []string{}
    for buffer.Scan(){
        stringArr = strings.Split(buffer.Text(), ",")
        sort.Strings(stringArr) 
    }
    
    result := 0
    for idx, val :=  range stringArr{
        sum := 0
        for _, c := range strings.ReplaceAll(val, "\"", "") {
            sum += int(c) - 64
        }
        result += sum * (idx + 1)
    }
    fmt.Println(result)
}
