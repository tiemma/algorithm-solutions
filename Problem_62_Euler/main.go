package main

import (
	"flag"
	"fmt"
	"math"
	"math/big"
	"sort"
	"strings"
)

type values struct {
	count  int
	number *big.Int
}

func main() {
	numPermutations := flag.Int("number", 5, "Number of permutations to find")
	flag.Parse()

	sets := map[string]*values{}
	for i := 1; i < math.MaxInt64; i++ {
		digitInt64 := int64(i)
		digit := big.NewInt(digitInt64)
		cube := big.NewInt(1)
		cube.Mul(cube, digit).Mul(cube, digit).Mul(cube, digit)

		cubeString := strings.Split(cube.String(), "")
		sort.Strings(cubeString)
		cubeStringSorted := strings.Join(cubeString, "")
		if _, ok := sets[cubeStringSorted]; !ok {
			sets[cubeStringSorted] = &values{number: cube, count: 0}
		}
		sets[cubeStringSorted].count += 1
		if sets[cubeStringSorted].count == *numPermutations {
			fmt.Println(sets[cubeStringSorted].number)
			return
		}
	}
}
