package main

import (
    "bufio"
    "flag"
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
)


func recurseAdd(arr [100][50]int, window int) []int{
    col := len(arr[0]) - 1
    result := []int{}
    rem := 0
    for col >= 0{
        temp := rem
        row := 0
        fmt.Println(row, " ", col, " ", temp)
        for row < len(arr) {
            temp += arr[row][col]
            row += 1               
        }
        result = append([]int{temp%10}, result...)
        fmt.Println(row, " ", col, " ", temp)
        rem = temp / 10
        fmt.Println(result)
        col -= 1
    }
    
    for rem > 0{
        result = append([]int{rem%10}, result...)
        rem /= 10
        fmt.Println(result)
    }
   // return result[len(result)-window:] 
    return result[:window] 
}



func main(){

    fileName := flag.String("file", "input.txt", "File to read integers from")
    window := flag.Int("window", 10, "Window of integers to  return")
    flag.Parse() 
     
    input, err := os.Open(*fileName)
    if err != nil {
        panic(err)
    }
    defer input.Close()

    buffer := bufio.NewScanner(input)
    stringArr := []string{}
    count := 0
    numArr := [100][50]int{}
    for buffer.Scan(){
        stringArr = strings.Split(buffer.Text(), "")
        for idx, el := range stringArr {
            numArr[count][idx], err = strconv.Atoi(el)
        }   
        count += 1 
    }
    var val float64 = 0 
    for idx, el := range recurseAdd(numArr, *window){
        val += float64(el) * math.Pow(10, float64(*window - idx - 1))
    } 
    fmt.Println("First ", *window, " digits are ", int(val)) 
}



