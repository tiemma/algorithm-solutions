package Linked_List_Cycle_2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "fmt"

func moveNode(head *ListNode) (bool, *ListNode) {
	if head.Next != nil {
		head = head.Next
		return true, head
	}
	return false, head
}

func getNodeAtIndex(head *ListNode, n int) *ListNode {
	for i := 0; i < n; i++ {
		head = head.Next
	}
	return head
}

func detectCycle(head *ListNode) *ListNode {
	one_step := head
	two_step := head

	end_two := false
	end_one := false

	count := 0

	for one_step != nil && two_step != nil{
		fmt.Println(count)
		fmt.Printf("node_1 %d == node_2 %d\n", one_step.Val, two_step.Val)

		end_one, one_step = moveNode(one_step)
		end_two, two_step = moveNode(two_step)
		end_two, two_step = moveNode(two_step)

		fmt.Printf("node_1 %d == node_2 %d\n\n", one_step.Val, two_step.Val)


		if two_step == one_step {
			if  end_two == false || end_one == false {
				return nil
			}

			return getNodeAtIndex(head, count)
		}

		count += 1

	}

	return nil
}