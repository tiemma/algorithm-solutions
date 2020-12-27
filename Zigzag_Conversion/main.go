package convert

// https://leetcode.com/problems/zigzag-conversion/

import (
    "strings"
)

func convert(s string, numRows int) string {
    if numRows == 1{
        return s
    }
    
    arr := make([]string, numRows)
    cord := 0
    factor := 1
    for _, val := range s{
        arr[cord] += string(val)
        if cord == numRows - 1 {
            factor = -1
        } else if cord == 0 {
            factor = 1
        }
        cord += factor
    }
    return strings.Join(arr, "")
    
}
