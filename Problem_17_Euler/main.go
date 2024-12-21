package main

//://projecteuler.net/problem=17 

import (
    "flag"
    "fmt"
    "math" 
    "strconv"
    "strings"
)



func getPowerExtensions(number int) (string, bool){
    extensions := map[int]string{}
    extensions[2] = "hundred"
    extensions[3] = "thousand"
    extensions[6] = "million"
    extensions[9] = "billion"

    if val, ok := extensions[number]; ok {
        return val, true
    }
    return "", false
} 

func getExtensions(number int) (string, bool){
    extensions := map[int]string{}
    extensions[1] = "one"
    extensions[2] = "two"
    extensions[3] = "three"
    extensions[4] = "four"
    extensions[5] = "five"
    extensions[6] = "six"
    extensions[7] = "seven"
    extensions[8] = "eight"
    extensions[9] = "nine"
    extensions[10] = "ten"
    extensions[11] = "eleven"
    extensions[12] = "twelve"
    extensions[13] = "thirteen"
    extensions[14] = "fourteen"
    extensions[15] = "fifteen"
    extensions[16] = "sixteen"
    extensions[17] = "seventeen"
    extensions[18] = "eighteen"
    extensions[19] = "nineteen"
    extensions[20] = "twenty"
    extensions[30] = "thirty"
    extensions[40] = "forty"  
    extensions[50] = "fifty"
    extensions[60] = "sixty"
    extensions[70] = "seventy"
    extensions[80] = "eighty"
    extensions[90] = "ninety"
    if val, ok := extensions[number]; ok {
        return val, true
    }
    return "", false
} 

func decomposePowers(power int) []float64 {
    result := []float64{}
    for power >= 3 {
        temp := (power / 3) * 3 
        result = append(result, float64(temp))
        power -= 3
    }
    return result
}


func getNumberIdentity(number int) []string {
    // For numbers below a thousand
    rsvPower := len(strconv.FormatInt(int64(number), 10)) - 1

    result := []string{}
    if val, ok := getExtensions(number); ok{
            result = append(result, val)
            rsvPower = -1
    }

    for rsvPower >= 0 {
        rsvPowerNum := int(math.Pow(10, float64(rsvPower)))
        temp, rem := number / rsvPowerNum, number % rsvPowerNum
        if val, ok := getExtensions(temp * rsvPowerNum); ok{
            result = append(result, val)
        } else {
            if val, ok := getExtensions(temp); ok{
                result = append(result, val)        
            }
        }
        if val, ok := getPowerExtensions(rsvPower); ok{       
            result = append(result, val)
            if rem > 0{
                result = append(result, "and") 
            }
        }
        if val, ok := getExtensions(rem); ok{
            result = append(result, val)
            rsvPower = 0
        }
        rsvPower -= 1
        number = rem
    }

    return result
}

func convertNumberToWords(number int) string{   
    var power float64 = float64(len(strconv.FormatInt(int64(number), 10)) - 1)  
    result := []string{}
    powerArr := decomposePowers(int(power))
    for _, power := range powerArr{
        divisor := math.Pow(10, power)
        temp, rem := int(float64(number)/divisor), number % int(divisor)
        
        result = append(result, getNumberIdentity(temp)...)
        if powerIdentifier, ok := getPowerExtensions(int(power)); ok{
            result = append(result, powerIdentifier)  
        }
        number = rem 
    }
    result = append(result, getNumberIdentity(number)...)        
    return strings.Join(result, " ")
}

 
func main(){
    number := flag.Int("number", 1_000, "Max number to sum to")
    flag.Parse()
    
    sum := 0
    for *number > 0{
        numInWords := convertNumberToWords(*number)
        numInWordsStripped := strings.Replace(numInWords, " ", "", -1)
        sum += len(numInWordsStripped)
        *number -= 1
    }    
    fmt.Println(sum)
}


