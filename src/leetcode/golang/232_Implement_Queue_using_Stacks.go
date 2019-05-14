package golang

import "github.com/golang-collections/collections/stack"

type MyQueue struct {
	input  *stack.Stack
	output *stack.Stack
	length int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		input:  stack.New(),
		output: stack.New(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.length++
	this.input.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.output.Len() == 0 {
		for this.input.Len() > 0 {
			this.output.Push(this.input.Pop())
		}
	}
	return this.output.Pop().(int)
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.output.Len() == 0 {
		for this.input.Len() > 0 {
			this.output.Push(this.input.Pop())
		}
	}
	return this.output.Peek().(int)
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.input.Len() == 0 && this.output.Len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
