package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//nums := []int{3, 2, 4}
	//fmt.Println(twoSum(nums, 6))

	//l1 := makeListNode(9)
	//l1.Next = makeListNode(9)
	//l2 := makeListNode(1)
	//res := addTwoNumbers(l1, l2)
	//fmt.Printf("%v", res)

	x := 1534236469
	fmt.Println(reverse(x))
}

// 计算2数之和，返回2个数字的下标
// [2, 7, 11, 8] 9 [0, 1]
func twoSum(nums []int, target int) []int {
	h := make(map[int]int)
	for k, v := range nums {
		h[v] = k
	}
	for k, v := range nums {
		expect := target - v
		if ek, ok := h[expect]; ok && ek != k {
			return []int{k, ek}
		}
	}
	return []int{}
}

func makeListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

/**
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		res       = new(ListNode)
		overfllow = 0
	)
	tmp := res
	m1, m2 := l1, l2
	for {
		sum := 0
		if m1 == nil && m2 != nil {
			sum = m2.Val + overfllow
		} else if m1 != nil && m2 == nil {
			sum = m1.Val + overfllow
		} else if m1 != nil && m2 != nil {
			sum = m1.Val + m2.Val + overfllow
		} else {
			break
		}
		tmp.Next = new(ListNode)
		overfllow = sum / 10
		tmp.Next.Val = sum % 10
		if m1 != nil {
			m1 = m1.Next
		}
		if m2 != nil {
			m2 = m2.Next
		}
		tmp = tmp.Next
	}
	if overfllow > 0 {
		tmp.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return res.Next
}

func reverse(x int) int {
	if x == 0 {
		return x
	}
	var res, i int
	i = x
	for math.Abs(float64(i)) >= 1 {
		res = res * 10 + i % 10
		i = i / 10
	}
	if res > 2<<30 || res < -2<<30 {
		return 0
	}
	return res
}