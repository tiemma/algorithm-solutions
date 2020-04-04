
# https://leetcode.com/problems/zigzag-conversion/

class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1:
            return s
        arr = {}
        for i in range(numRows):
            arr[i] = ""
        cord = 0
        const = 1
        reset = True
        for string in s:
            arr[cord] += string
            if cord == numRows - 1:
                const = -1
            elif cord == 0:
                const = 1
            cord += const
        result = ""
        for i in arr:
            result += arr[i]
        return result


                
                
                
        
