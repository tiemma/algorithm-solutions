package main


import (
    "fmt"
    "flag"
    "math"
    
)


func getDivisorSum(num int) int{
    i := 2
    sum := 0
    for i <= int(math.Sqrt(float64(num))) {
        if num % i == 0 {
            if i != num/i {
                sum += i + num / i 
            } else {
                sum += i
            }
        }
        i += 1
    }  
    return sum + 1
}

func isUnique(number int, result map[int]bool) bool {
    for k, _ := range result {
        if _, ok := result[int(math.Abs(float64(number -k)))]; ok{
            return false      
        }

    }
    return true
}

func main(){
    number := flag.Int("number", 28123, "Range to detect sums over")
    flag.Parse()
    
    /*
    Source: https://mathworld.wolfram.com/AbundantNumber.html
    Also, as opposed to the limit on the dual abundant number sum on ProjectEuler, the actual limit is 20161 
    For abundant numbers,
    1. Multiples of perfect numbers or abundant number are abundant
    2. Primes are not abundant

    So for this solution, we can also apply an elimination method instead of a full blown search to remove numbers as such:
    1. Prime numbers are deficient
    2. Powers of primes are deficient
    3. Perfect numbers can be excluded 
    4. Perfect number divisors are also deficient
    5. Deficient number divisors are also deficient
    */
   
    result := map[int]bool{}
    i := 2 
    sum := *number * (*number + 1) / 2

    for i <= *number{
        if i < getDivisorSum(i) {
            result[i] = true
        }
        if !isUnique(i, result){
            sum -= i
            fmt.Println("Number is valid" , i)
        } 

        i += 1
    }
    fmt.Println(sum)
}
