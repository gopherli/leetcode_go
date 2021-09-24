package day01

import "testing"

func TestOffer30(t *testing.T) {
	minStack := Constructor30()
	minStack.Push(2)
	minStack.Push(0)
	minStack.Push(3)
	minStack.Push(0)
	minStack.Min()
	minStack.Pop()
	minStack.Min()
	minStack.Pop()
	minStack.Min()
	minStack.Pop()
	minStack.Min()
}
