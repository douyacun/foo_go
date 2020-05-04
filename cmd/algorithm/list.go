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
	//a := []int{1, 2, 3}
	//list := makeList(a)
	//
	//list = removeNthFromEnd(list, 2)
	l1 := makeList([]int{1, 2})
	l2 := makeList([]int{1, 3, 5})
	list := mergeTwoLists(l1, l2)
	fmt.Printf("%v", list)
}

// 尾插法
func makeList(arr []int) *ListNode {
	var list = new(ListNode)
	tmp := list
	for i := 0; i < len(arr); i++ {
		node := &ListNode{
			Val:  arr[i],
			Next: nil,
		}
		tmp.Next = node
		tmp = tmp.Next
	}
	return list.Next
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
		res = res*10 + i%10
		i = i / 10
	}
	if res > 2<<30 || res < -2<<30 {
		return 0
	}
	return res
}

/**
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

思路：快慢指针
快指针和慢指针间隔n+1
快指针到达尾部func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		i, j = head, head
	)
	var k = 1
	for j.Next != nil && k < n+1 {
		j = j.Next
		k++
	}
	if k < n {
		return head
	} else if k == n {
		return head.Next
	}
	for j.Next != nil {
		j = j.Next
		i = i.Next
	}
	i.Next = i.Next.Next
	return head
}
慢指针.Next = 慢指针.Next.Next
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		i, j = head, head
	)
	var k = 1
	for j.Next != nil && k < n+1 {
		j = j.Next
		k++
	}
	if k < n {
		return head
	} else if k == n {
		return head.Next
	}
	for j.Next != nil {
		j = j.Next
		i = i.Next
	}
	i.Next = i.Next.Next
	return head
}

/**
Definition for singly-linked list.
type ListNode struct {
     Val int
     Next *ListNode
}

将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = new(ListNode)
	tmp := head

	for l1 != nil {
		for l2 != nil && l2.Val <= l1.Val {
			tmp.Next = l2
			l2 = l2.Next
			tmp = tmp.Next
		}
		tmp.Next = l1
		l1 = l1.Next
		tmp = tmp.Next
	}
	if l2 != nil {
		tmp.Next = l2
	}
	return head.Next
}
