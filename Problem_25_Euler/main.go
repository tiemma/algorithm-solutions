package main

import (
    "fmt"
    "flag"
    "image/color"
    "math"
    "math/big"

    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/vg"
)

<<<<<<< HEAD
// // https://projecteuler.net/problem=25
=======

>>>>>>> 824725c56150b0a5d9994b1f9f99dc0863490bbf

func getDigitLengthBinet(n int) int {
     // We use a reverse implementation of the Binet formula
    // Binet Formula gives the fibonacci digit but we can apply 
    // A log to that to that formula to obtain the digit count
   if n < 2 {
        return 1
    }
    
    var phi float64= (1 + math.Sqrt(5)) / 2
    
    nDigits := float64(n) * math.Log10(phi) - (math.Log10(5) / 2) 
    
    return int(math.Ceil(nDigits))
}


func plotDifference(num float64){
    p, err := plot.New()
    if err != nil {
        panic(err)
    }
    p.Title.Text = "Differences"
    p.X.Label.Text = "X"
    p.Y.Label.Text = "Y"
   
    exp := plotter.NewFunction(func(x float64) float64 {  
        diff := float64(getFnIndex(int(x)) - getFnIndex(int(x - 1)))
        return diff 
    })
    exp.Dashes = []vg.Length{vg.Points(4), vg.Points(4)}
    exp.Width = vg.Points(2)
    exp.Color = color.RGBA{G: 255, A: 255}

       // Add the functions and their legend entries.
    p.Add(exp)
    p.Legend.Add("len(F(n)) - len(F(n-1))", exp)
    p.Legend.ThumbnailWidth = 2 * vg.Inch

    // Set the axis ranges.  Unlike other data sets,
    // functions don't set the axis ranges automatically
    // since functions don't necessarily have a
    // finite range of x and y values.
    p.X.Min = 0
    p.X.Max = num
    p.Y.Min = 0
    p.Y.Max = 10 

    // Save the plot to a PNG file.
    if err := p.Save(16*vg.Inch, 16*vg.Inch, fmt.Sprintf("differences-%1f.png", num)); err != nil { 
        panic(err)
    }
}



func plotFunc(num float64){
    p, err := plot.New()
    if err != nil {
        panic(err)
    }
    p.Title.Text = "Functions"
    p.X.Label.Text = "X(5*k)"
    p.Y.Label.Text = "Y(k)"

    quad := plotter.NewFunction(func(x float64) float64 { return x / 5 })
    quad.Color = color.RGBA{B: 255, A: 255}

   
    exp := plotter.NewFunction(func(x float64) float64 { return float64(getDigitLengthBinet(int(x))) })
    exp.Dashes = []vg.Length{vg.Points(1), vg.Points(1)}
    exp.Width = vg.Points(2)
    exp.Color = color.RGBA{G: 255, A: 255}

   
        sin := plotter.NewFunction(func(x float64) float64 { return - (x/5) + float64(getDigitLengthBinet(int(x)))  })
    sin.Dashes = []vg.Length{vg.Points(4), vg.Points(5)}
    sin.Width = vg.Points(4)
    sin.Color = color.RGBA{R: 255, A: 255}

     // Add the functions and their legend entries.
    p.Add(quad, exp, sin)
    p.Legend.Add("5 * k", quad)
    p.Legend.Add("getDigitLengthBinet", exp)
    p.Legend.Add("getDigitLengthBinet - 5 * k", sin)
    p.Legend.ThumbnailWidth = 2 * vg.Inch
    p.Legend.Top = true

    // Set the axis ranges.  Unlike other data sets,
    // functions don't set the axis ranges automatically
    // since functions don't necessarily have a
    // finite range of x and y values.
    p.X.Min = 0
    p.X.Max = num
    p.Y.Min = 0
    p.Y.Max = num * 1 / 4 

    // Save the plot to a PNG file.
    if err := p.Save(16*vg.Inch, 16*vg.Inch, fmt.Sprintf("functions-%1f.png", num)); err != nil { 
        panic(err)
    }
}


func getDigitLength(n int) int{
   
	
    bigFive := big.NewFloat(5)
    sqrt5 := big.NewFloat(1).Sqrt(bigFive)
    phi := big.NewFloat(1)

    phi.Quo(phi.Add(big.NewFloat(1), sqrt5), big.NewFloat(2))
  
    temp := n
    val := big.NewFloat(1)
        // Raise val to the power of n
    // math/big doesn't have Exp on Float 
    // So I just use a simple loop to multiply itself
    for temp > 0{
        val = val.Mul(val, phi)
        temp -= 1
    }
    
    ans := val.Quo(val, sqrt5)
	
    
    fmt.Printf("Value at n: %f in fibonnaci sequence is: %f\n", n, ans)
    // Go returns the exp in powers of 2
    // 2 ^^ y == 10 ^^ x
    // Taking log
    // y log 2 = x log 10 = x 
    // x in base10 = y log 2 hence the addition of that factor
    return int(float64(ans.MantExp(ans)) * math.Log10(2)) + 1
}

func getFnIndex(numberOfDigits int) int{
    // From my analysis, the fibonacci number index follows a plot along 5x of the number
    // approximately in relation to its number of digits 
    maxDelta := 5 * (numberOfDigits)
   
    for maxDelta > 0 {
        if numberOfDigits > getDigitLengthBinet(maxDelta) {
            fmt.Println("Approximate value", 5 * numberOfDigits)
            fmt.Println("Actual value", maxDelta + 1)
            fmt.Println("Offset iterations", 5 * numberOfDigits  - (maxDelta + 1))
            fmt.Println("Iteration Percentage", float64((5 * numberOfDigits)  - (maxDelta + 1)) * 100 / float64(maxDelta + 1))
            break
        } 
        maxDelta --
    }
    return maxDelta + 1
}

func main(){
    numberOfDigits := flag.Int("number", 2, "Number of digits to look out for")
    plotGraphs := flag.Bool("plot", false, "Boolean to plot graph of deviations")
    flag.Parse()

    
    maxDelta := 5 * *numberOfDigits 
    if *plotGraphs { 
        plotFunc(float64(maxDelta))
        plotDifference(float64(maxDelta))
    }
   
     
    i := 2
    result := []int{}
    prevDigitCount := 1
    prevIndex := 0
    for i < *numberOfDigits{
        if prevDigitCount < getDigitLengthBinet(i) {
            result = append(result, i - prevIndex)
            prevIndex = i 
            prevDigitCount = getDigitLengthBinet(i)
        }
        i += 1
    }
    fmt.Println(result)
    // getDigitLength(*numberOfDigits)
    
    getFnIndex(*numberOfDigits)
}
