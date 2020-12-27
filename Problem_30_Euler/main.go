package main

import (
  "flag"
  "fmt"
  "math"
)


func getDigitPowerSum(num float64, power float64) float64{
  sum := *new(float64)
  for num > 0{
    sum += math.Pow(math.Mod(num, 10), power)
    num = math.Floor(num / 10)
  }
  return sum
}

func main(){
  power := flag.Float64("power", 5, "Power to raise individual digits to")
  flag.Parse()

  
  limit := math.Pow(9, *power) * *power 
  sum := *new(float64)
  for limit > 9 {
    if limit == getDigitPowerSum(limit, *power) {
      sum += limit
    }
    limit --
  }
  fmt.Println(sum) 
}

