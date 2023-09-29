package leetcode

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slowPointer := head
	fastPointer := head

	//Making mistake by not checking slowPointer.Next != nil && fastPointer.Next != nil && fastPointer.Next.Next != nil
	for head != nil && slowPointer.Next != nil && fastPointer.Next != nil && fastPointer.Next.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next

		if slowPointer == fastPointer {
			return true
		}
		head = head.Next
	}

	return false
}
