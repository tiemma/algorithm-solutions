package Single_Cycle_Check


// https://www.algoexpert.io/questions/Single%20Cycle%20Check

func modLikePython(d int, m int) int {
	var res int = d % m
	if ((res < 0 && m > 0) || (res > 0 && m < 0)) {
		return res + m
	}
	return res
}


// O(n) time and space
func HasSingleCycle(array []int) bool {
	// Write your code here.
	visited := make([]bool, len(array))
	curr_idx := 0
	num_el_visited := 0
	for num_el_visited < len(array){
		curr_idx = modLikePython(curr_idx + array[curr_idx],  len(array))
		if visited[curr_idx]{
			return false
		}
		visited[curr_idx] = true
		num_el_visited += 1
	}
	return curr_idx == 0
}


// Optimised Version
// O(n) time
// O(1) space
func HasSingleCycle(array []int) bool {
	curr_idx := 0
	num_el_visited := 0
	for num_el_visited < len(array){
		if num_el_visited > 0 && curr_idx == 0{
			return false
		}
		curr_idx = modLikePython(curr_idx + array[curr_idx],  len(array))
		num_el_visited += 1
	}
	return curr_idx == 0
}