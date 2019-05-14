package golang

type MyNode struct {
	Val  int
	Prev *MyNode
	Next *MyNode
}

type MyLinkedList struct {
	Head *MyNode
	Tail *MyNode
	size int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) GetNode(index int) *MyNode {
	if index < 0 || index >= this.size {
		return nil
	}
	curr := this.Head
	for ; index > 0; index-- {
		curr = curr.Next
	}
	return curr
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	curr := this.GetNode(index)
	if curr != nil {
		return curr.Val
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	head := this.Head
	this.Head = &MyNode{
		Val:  val,
		Prev: nil,
		Next: head,
	}
	if head != nil {
		head.Prev = this.Head
	} else {
		this.Tail = this.Head
	}
	this.size += 1
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	tail := this.Tail
	this.Tail = &MyNode{
		Val:  val,
		Prev: tail,
		Next: nil,
	}
	if tail != nil {
		tail.Next = this.Tail
	} else {
		this.Head = this.Tail
	}
	this.size += 1
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
	} else if index == this.size {
		this.AddAtTail(val)
	} else if index > this.size {
		return
	} else {
		next := this.GetNode(index)
		prev := next.Prev
		curr := MyNode{
			Val:  val,
			Prev: prev,
			Next: next,
		}
		prev.Next = &curr
		next.Prev = &curr
		this.size += 1
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	curr := this.GetNode(index)
	prev := curr.Prev
	next := curr.Next
	if prev == nil {
		this.Head = next
	} else {
		prev.Next = next
	}
	if next == nil {
		this.Tail = prev
	} else {
		next.Prev = prev
	}
	this.size -= 1
}
