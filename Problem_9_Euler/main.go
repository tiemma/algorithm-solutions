package main

//https://projecteuler.net/problem=9


import (
    "fmt"
    "math"
    "flag"
)

func BinetFormula(n float64) float64 {
	sqrt5 := math.Sqrt(5)
	phi := (1 + sqrt5) / 2
	ans := math.Round(math.Pow(phi, n) / sqrt5)
	fmt.Printf("Value at n: %f in fibonnaci sequence is: %f\n", n, ans)
	return ans
}

func getTriple(n float64) [3]float64{
    firstTriple := BinetFormula(n) * BinetFormula(n+3)
    secondTriple := BinetFormula(n+1) * BinetFormula(n+2) * 2
    thirdTriple := math.Pow(BinetFormula(n+1), 2) + math.Pow(BinetFormula(n+2), 2) 
    return [...]float64{firstTriple, secondTriple, thirdTriple}
}


func validateConditionals(maxSum float64, firstTriple float64, secondTriple float64, thirdTriple float64) bool{
    return math.Pow(firstTriple, 2) + math.Pow(secondTriple, 2) == math.Pow(thirdTriple, 2) && firstTriple + secondTriple + thirdTriple == maxSum 

}

func getPythagoreanTriple(maxSum float64) [3]float64 {
    var n float64 = 0
    for n < math.Sqrt(maxSum){        
        triples := getTriple(n)
        firstTriple := triples[0]
        secondTriple :=  triples[1]
        thirdTriple :=  triples[2] 
        if validateConditionals(maxSum, firstTriple, secondTriple, thirdTriple) {
            return getTriple(n)      
        }
        n++
    }
    return [...]float64{0, 0, 0}
}

func getPythagoreanTripleSquared(maxSum float64) [3]float64 {
    var n float64 = 0
    for n < math.Sqrt(maxSum){        
        firstTriple := 2 * n
        secondTriple :=  math.Pow(n, 2) - 1
        thirdTriple := math.Pow(n, 2) + 1
        if validateConditionals(maxSum, firstTriple, secondTriple, thirdTriple) {
            return [...]float64 {firstTriple, secondTriple, thirdTriple} 
        }
        n++
    }
    return [...]float64{0, 0, 0}
}


func getPythagoreanTripleBruteForce(n float64) [3]float64{
   var i float64 = 1
   var j float64  = 0
    var k float64 = 0
   for i < n {
        j = 0
        for j < n {
            k = n - i - j
                 
            // fmt.Println("Using the following values: ", i, j, k, n) 
            if validateConditionals(n, i, j, k) {
                return [...]float64 {i, j, k} 
            }
            j++
        }
        i++
    } 
    return [...]float64{0, 0, 0}
}


func main(){
    maxNumPtr := flag.Float64("number", 1000, "Integer for the max range")
    flag.Parse()
    solution := getPythagoreanTripleBruteForce(*maxNumPtr)

    fmt.Println("Product array is ", solution)
    fmt.Println("Answer is: ", int(solution[0] * solution[1] * solution[2]))
}


