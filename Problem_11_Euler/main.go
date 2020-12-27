package main


// https://projecteuler.net/problem=11

import (
    "bufio"
    "flag"
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
)


func upperDiagonal(arr [20][20]int, row int, col int, grid int, window int) int {
    j := 0
    product := 1
    if row <= grid - window && col >= window - 1 {
       for j < window {
            product *= arr[row+j][col-j]                
            j += 1   
        }
    }
    fmt.Println(row, col, product)
    return product
}



func lowerDiagonal(arr [20][20]int, row int, col int, grid int, window int) int {
    j := 0
    product := 1
    if row <= grid - window && col <= grid - window {
       for j < window {
            product *= arr[row+j][col+j]                
            j += 1   
        }
    }
    return product
}

func up(arr [20][20]int, row int, col int, grid int, window int) int {
    j := 0
    product := 1
    if row >= window - 1 {
       for j < window {
            product *= arr[row-j][col]                
            j += 1   
        }
    }
    return product
}

func down(arr [20][20]int, row int, col int, grid int, window int) int {
    j := 0
    product := 1
    if row <= grid - window {
       for j < window {
            product *= arr[row+j][col]                
            j += 1   
        }
    }
    return product
}

func left(arr [20][20]int, row int, col int, grid int, window int) int {
    j := 0
    product := 1
    if col >= window - 1 {
       for j < window {
            product *= arr[row][col-j]                
            j += 1   
        }
    }
    return product
}

func right(arr [20][20]int, row int, col int, grid int, window int) int {
    j := 0
    product := 1
    if col <= grid - window {
       for j < window {
            product *= arr[row][col+j]                
            j += 1   
        }
    }
    return product
}



func setMaxProduct(product int, temp int) int{
    if temp > product {
        product = temp
    }
    return product
} 

func maxDiagonalProduct(arr [20][20]int, window int, grid int) int{
    i := 0
    product := math.MinInt64
    for i < int(len(arr) * len(arr[0])) {
        row, col := i / grid, i % grid
        product = setMaxProduct(product, down(arr, row, col, grid, window))  
        product = setMaxProduct(product, up(arr, row, col, grid, window))  
        product = setMaxProduct(product, right(arr, row, col, grid, window))  
        product = setMaxProduct(product, lowerDiagonal(arr, row, col, grid, window))  
        product = setMaxProduct(product, upperDiagonal(arr, row, col, grid, window))  
        product = setMaxProduct(product, left(arr, row, col, grid, window))  
        i += 1
    }
    return product 
}   


func main(){

    fileName := flag.String("file", "input.txt", "File to read integers from")
    window := flag.Int("window", 4, "Window along the diagonal to multiply")
    flag.Parse() 
     
    input, err := os.Open(*fileName)
    if err != nil {
        panic(err)
    }
    defer input.Close()

    buffer := bufio.NewScanner(input)
    stringArr := []string{}
    grid := 20
    numArr := [20][20]int{}
    for buffer.Scan(){
        stringArr = strings.Split(buffer.Text(), " ")
        for idx, el := range stringArr {
            numArr[idx/grid][idx%grid], err = strconv.Atoi(el) 
        }   
    }
    fmt.Println(maxDiagonalProduct(numArr, *window, grid))
}

