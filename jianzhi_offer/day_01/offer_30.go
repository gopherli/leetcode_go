package day01

// Reference：https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/

type MinStack struct {
	min  []int
	nums []int
}

/** initialize your data structure here. */
func Constructor30() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.nums = append(this.nums, x)
	if len(this.min) == 0 {
		this.min = append(this.min, x)
	} else if this.min[len(this.min)-1] > x {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, this.min[len(this.min)-1])
	}
}

func (this *MinStack) Pop() {
	this.nums = this.nums[:len(this.nums)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) Min() int {
	return this.min[len(this.min)-1]
}
