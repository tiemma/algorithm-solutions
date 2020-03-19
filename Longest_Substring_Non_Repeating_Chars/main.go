func find(arr []string, v string) (bool, int) {
    for idx, str := range arr{
        if v == str{
            return true, idx
        }
    }
    return false, -1
}

func lengthOfLongestSubstring(s string) int {
    ans := 0
    result := []string{}
    
    for _, v := range s{
        found, idx := find(result, string(v))
        if found {
            result = result[idx+1:]
        }
        result = append(result, string(v))
        if ans < len(result) {
            ans = len(result)
 :       } 
        
    }
    return ans
}
