package main

// https://projecteuler.net/problem=4

// 998001   999 x 999

// 996699

// 100000   100 x 100


import (
    "math"
    "fmt"
)

func reverseNumber(n int) int{
    reverse := 0
    for n > 0 {
        lastDigit := n % 10
        reverse = (reverse * 10) + lastDigit
        n = n / 10
    }
    return reverse
}


func getLargestPalindrome() int{
    maxPalindrome := -math.MaxInt64
    for i :=1000; i > 0; i-- {
	    for j := 1000; j > 0; j-- {
            product := i * j
            if product  ==  reverseNumber(product) && product > maxPalindrome {
                maxPalindrome = product
            }
        }
    }
    return maxPalindrome
}


func main(){
    fmt.Println(getLargestPalindrome())
}
