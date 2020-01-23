package Depth_First_Search


// https://www.algoexpert.io/questions/Depth-first%20Search



// Do not edit the class below except
// for the depthFirstSearch method.
// Feel free to add new properties
// and methods to the class.
type Node struct {
	Name     string
	Children []*Node
}


// Optimised Version O(V + E)
// V - Vertex
// E - Edges
func (n *Node) DepthFirstSearch(array []string) []string {
	// Write your code here.
	array = append(array, n.Name)
	for _, node := range n.Children {
		array = node.DepthFirstSearch(array)
	}

	return array
}
