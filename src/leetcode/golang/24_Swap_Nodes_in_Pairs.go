package golang

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {
	var guard, prev, curr, next *ListNode
	guard = &ListNode{Next: head}
	prev = guard
	for prev.Next != nil && prev.Next.Next != nil {
		curr = prev.Next
		next = curr.Next
		prev.Next, curr.Next, next.Next = next, next.Next, curr
		prev = curr
	}
	return guard.Next
}
