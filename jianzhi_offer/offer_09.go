package jianzhi_offer

import (
	"fmt"
	"time"
)

// Reference：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
type CQueue struct {
	in  stack
	out stack
}

type stack []int

func (s *stack) Push(value int) {
	start := time.Now().Nanosecond()
	*s = append(*s, value)
	end := time.Now().Nanosecond()
	fmt.Println("Push speed time => ", end-start)
}

func (s *stack) Pop() int {
	start := time.Now().Nanosecond()
	top := len(*s)
	value := (*s)[0]
	*s = (*s)[1:top]
	end := time.Now().Nanosecond()
	fmt.Println("Pop speed time => ", end-start)
	return value
}

func Constructor09() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.in.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.out) == 0 {
		if len(this.in) > 0 {
			return this.in.Pop()
		}
	} else {
		this.in.Pop()
		return this.out.Pop()
	}
	return -1
}
