package Kadane_Algorithm


//algoexpert.io/questions/Kadane's%20Algorithm

func max(a int, b int) int{
	if a > b {
		return a
	}
	return b
}

// Optimised version
// O(n) time
// O(1) space
func KadanesAlgorithm(array []int) int {
	max_curr, max_at_idx := array[0], array[0]
	for idx := 1; idx < len(array); idx++{
		el := array[idx]
		max_at_idx = max(el, max_at_idx + el)
		max_curr = max(max_curr, max_at_idx)
	}
	return max_curr
}
