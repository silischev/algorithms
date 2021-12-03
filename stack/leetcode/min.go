package main

type MinStack struct {
	top  int
	min  int
	data []int
}

func NewMinStack() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(val int) {
	if len(ms.data) == 0 || ms.min > val {
		ms.min = val
	}

	ms.top = val
	ms.data = append(ms.data, val)
}

func (ms *MinStack) Pop() {
	if len(ms.data) == 0 {
		return
	}

	ms.data = ms.data[:len(ms.data)-1]

	if len(ms.data) == 0 {
		return
	}

	ms.top = ms.data[len(ms.data)-1]

	min := ms.data[0]
	for _, val := range ms.data {
		if val < min {
			min = val
		}
	}

	ms.min = min
}

func (ms *MinStack) Top() int {
	return ms.top
}

func (ms *MinStack) GetMin() int {
	return ms.min
}
