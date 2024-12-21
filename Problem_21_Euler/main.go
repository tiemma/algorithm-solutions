package main


// https://projecteuler.net/problem=21
import (
    "fmt"
    "flag"
    "math"
)
func getMultipleSum(num int) int{
    result := []int{}
    sum := 0
    i := 2
    for i < int(math.Sqrt(float64(num))) {
        if num % i == 0 {
            result = append(result, num / i, i)
        }  
        i++ 
    }
      for _, val:= range result{
        sum += val
    }
    return sum
}

func main(){
    maxNum := flag.Int("window", 10_000, "Limit to check for amicable pairs from 0")
    flag.Parse()

    usedValues :=map[int]bool{}
    result := 0 
    
    for *maxNum >= 2 {
        temp := 1 + getMultipleSum(*maxNum) 
        if 1 + getMultipleSum(temp) == *maxNum && temp != *maxNum {
            if _, ok := usedValues[temp]; !ok{
                if _, ok := usedValues[*maxNum]; !ok {
                    usedValues[temp] = true
                    usedValues[*maxNum] = true
                    fmt.Println(temp, *maxNum)
                    result += temp + *maxNum
                }
            }
        }
        *maxNum--
    }
    fmt.Println(result)
}
