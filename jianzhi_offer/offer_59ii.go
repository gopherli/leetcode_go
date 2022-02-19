package jianzhi_offer

// https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof/

type MaxQueue struct {
	queue []int
	max   []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (m *MaxQueue) Max_value() int {
	if len(m.max) == 0 {
		return -1
	}
	return m.max[0]
}

func (m *MaxQueue) Push_back(value int) {
	m.queue = append(m.queue, value)
	for len(m.max) > 0 && value > m.max[len(m.max)-1] {
		m.max = m.max[:len(m.max)-1]
	}
	m.max = append(m.max, value)
}

func (m *MaxQueue) Pop_front() int {
	if len(m.queue) == 0 {
		return -1
	}
	pop := m.queue[0]
	m.queue = m.queue[1:]
	if pop == m.max[0] {
		m.max = m.max[1:]
	}
	return pop
}
