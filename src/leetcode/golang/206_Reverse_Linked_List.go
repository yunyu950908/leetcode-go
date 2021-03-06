package golang

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev, curr *ListNode
	curr = head
	for curr != nil {
		prev, curr, curr.Next = curr, curr.Next, prev
	}
	return prev
}
