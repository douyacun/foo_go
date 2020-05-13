package main

import "fmt"

type KthLargest struct {
	q    []int
	size int
	max  int
}

func main() {
	nums := []int{4, 5, 8, 2}
	obj := Constructor(3, nums)
	k := obj.Add(1)
	k = obj.Add(9)
	k = obj.Add(100)
	fmt.Println(k)
}

func Constructor(k int, nums []int) KthLargest {
	h := KthLargest{
		q:    make([]int, k+1),
		size: 0,
		max:  k,
	}
	for _, v := range nums {
		h.Add(v)
	}
	return h
}

func (h *KthLargest) Add(val int) int {
	if h.size < h.max {
		h.size++
		h.q[h.size] = val
		h.up(h.size)
	} else {
		if h.q[1] > val {
			return h.q[1]
		} else {
			h.q[1] = val
			h.down(1)
		}
	}
	return h.q[1]
}

func (h *KthLargest) up(i int) {
	for i > 1 && h.less(i, h.parent(i)) {
		h.q[h.parent(i)], h.q[i] = h.q[i], h.q[h.parent(i)]
		i = h.parent(i)
	}
}

func (h *KthLargest) down(i int) {
	for h.left(i) <= i {
		min := h.left(i)
		if h.right(i) <= h.max && h.less(h.right(i), h.left(i)) {
			min = h.right(i)
		}
		if h.less(i, min) {
			break
		}
		h.q[i], h.q[min] = h.q[min], h.q[i]
		i = min
	}
}

func (h *KthLargest) less(first, second int) bool {
	return h.q[first] < second
}

func (h *KthLargest) left(i int) int {
	return 2 * i
}

func (h *KthLargest) right(i int) int {
	return 2*i + 1
}

func (h *KthLargest) parent(i int) int {
	return i / 2
}
