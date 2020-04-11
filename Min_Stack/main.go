package main

import (
	"math"
)

type MinStack struct {
	stack []int
	size  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	this.size += 1
}

func (this *MinStack) Pop() {
	if this.size > 0 {
		this.size -= 1
		this.stack = this.stack[:this.size]
	}
}

func (this *MinStack) Top() int {
	return this.stack[this.size-1]
}

func (this *MinStack) GetMin() int {
	min := math.MaxInt64
	for _, elem := range this.stack {
		if elem < min {
			min = elem
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
  * obj := Constructor();
	 * obj.Push(x);
	  * obj.Pop();
		 * param_3 := obj.Top();
		  * param_4 := obj.GetMin();
*/
