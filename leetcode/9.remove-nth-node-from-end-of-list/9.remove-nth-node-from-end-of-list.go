/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dummyNode := &ListNode{Val: -1, Next: head}
	l := dummyNode
	r := head
	for n > 0 {
		r = r.Next
		n = n - 1
	}

	for r != nil {
		l = l.Next
		r = r.Next
	}

	//remove the nth node, l.Next node
	//nxt := l.Next.Next
	l.Next = l.Next.Next

	return dummyNode.Next
}