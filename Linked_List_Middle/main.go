package main

// https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/529/week-2/3290/

/**
 * Definition for singly-linked list.
  * type ListNode struct {
		 *     Val int
		  *     Next *ListNode
			 * }
*/
func middleNode(head *ListNode) *ListNode {
	prev_head := head
	end := false
	for {
		if head == nil {
			return prev_head
		}
		if head.Next != nil {
			head = head.Next.Next
		} else {
			end = true
		}
		if end {
			return prev_head
		}
		prev_head = prev_head.Next
	}
}
