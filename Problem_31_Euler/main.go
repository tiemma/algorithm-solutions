package main


import (
  "flag"
  "fmt"
)

func main(){
  amount := flag.Int("amount", 200, "Amount to get from permutative sums")
  flag.Parse()  

  coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
  sum := make([]int, *amount)
  sum = append([]int{1}, sum...)
  for _, val := range coins {
    for j := val; j <= *amount; j++ {
      sum[j] += sum[j - val]
    } 
  }
  fmt.Println(sum[*amount])

}
