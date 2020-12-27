package Move_Element_To_End


// https://www.algoexpert.io/questions/Move%20Element%20To%20End


// Optimised version O(n)
func MoveElementToEnd(array []int, toMove int) []int {
	lastIdx := 0
	for idx := 1; idx < len(array); idx++{
		if(array[lastIdx] == toMove && array[idx] != toMove){
			array[idx], array[lastIdx] = array[lastIdx], array[idx]
			lastIdx += 1
		} else if(array[lastIdx] != toMove){
			lastIdx += 1
		}
	}
	return array
}
