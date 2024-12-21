package main

// https://projecteuler.net/problem=29

import (
  "flag"
  "fmt"
  "math/big"
)

func main() {
  a := flag.Int("a", 100, "Window for power combination as value")
  b := flag.Int("b", 100, "Window for power combination as power")

  flag.Parse()
  mem := map[string]bool{}
  count := 2
  for *a >= count  {
    temp := 2
    for *b >= temp {
      val := big.NewInt(1)
      val.Exp(big.NewInt(int64(count)), big.NewInt(int64(temp)), nil)
      mem[val.String()] = true
      temp ++
    }
     count ++
  }
  fmt.Println(len(mem))
}
