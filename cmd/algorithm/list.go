package main

import (
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
	//l1 := makeList([]int{1, 2, 6})
	//l2 := makeList([]int{1, 4, 5, 6})
	//l3 := makeList([]int{1, 3, 9})
	//list := mergeKLists([]*ListNode{l1, l2, l3})
	//fmt.Printf("%v", list)

	//l2 := makeList([]int{1, 2, 3, 4, 5, 6, 7, 8})
	//list := swapPairs2(l2)
	//fmt.Sprintf("%v", list)

	//l2 := makeList([]int{1, 2, 3, 4, 5, 6, 7, 8})
	//list := reverseKGroup(l2, 3)
	//fmt.Sprintf("%v", list)

	l2 := makeList([]int{1, 2, 3, 4, 5})
	list := rotateRight(l2, 2)
	printList(list)
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
* 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
* 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
* 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
*
* 示例：
*
* 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
* 输出：7 -> 0 -> 8
* 原因：342 + 465 = 807
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
* 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
*
* 示例：
*
* 给定一个链表: 1->2->3->4->5, 和 n = 2.
*
* 当删除了倒数第二个节点后，链表变为 1->2->3->5.
* 说明：
*
* 给定的 n 保证是有效的。
*
* 进阶：
*
* 你能尝试使用一趟扫描实现吗？
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
* Definition for singly-linked list.
* type ListNode struct {
*      Val int
*      Next *ListNode
* }
*
* 将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
*
* 示例：
*
* 输入：1->2->4, 1->3->4
* 输出：1->1->2->3->4->4
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	header := new(ListNode)
	temp := header

	for l1 != nil {
		for l2 != nil && l2.Val <= l1.Val {
			temp.Next = l2
			l2 = l2.Next
			temp = temp.Next
		}
		temp.Next = l1
		l1 = l1.Next
		temp = temp.Next
	}
	if l2 != nil {
		temp.Next = l2
	}
	return header.Next
}

/**
* 合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
* 1. 顺序合并
* 2. 分治合并
* 3. 最小堆
 */
func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}
func merge(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	left := merge(lists, start, mid)
	right := merge(lists, mid+1, end)
	return mergeTwoLists(left, right)
}

/**
* 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
* 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
* 给定 1->2->3->4, 你应该返回 2->1->4->3.
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	firstNode := head
	secondNode := head.Next

	firstNode.Next = swapPairs(secondNode.Next)
	secondNode.Next = firstNode

	return secondNode
}

func swapPairs2(head *ListNode) *ListNode {
	header := &ListNode{
		Val:  0,
		Next: head,
	}
	prev := header
	for prev.Next != nil && prev.Next.Next != nil {
		prev.Next.Next, prev.Next.Next.Next, prev.Next = prev.Next.Next.Next, prev.Next, prev.Next.Next
		prev = prev.Next.Next
	}
	return header.Next
}

/**
* 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
* k 是一个正整数，它的值小于或等于链表的长度。
* 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
*
* 示例：
* 给你这个链表：1->2->3->4->5
* 当 k = 2 时，应当返回: 2->1->4->3->5
* 当 k = 3 时，应当返回: 3->2->1->4->5
* 说明：
* 你的算法只能使用常数的额外空间。
* 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	node := &ListNode{
		Val:  0,
		Next: head,
	}
	prev := node
	tail := node
	for true {
		c := 1
		for tail != nil && c <= k {
			tail = tail.Next
			c++
		}
		if tail == nil {
			break
		}
		nextHead := prev.Next
		var cur *ListNode
		for prev.Next != tail {
			// 拆出cur
			cur, prev.Next = prev.Next, prev.Next.Next
			// cur拼到tail后面
			cur.Next, tail.Next = tail.Next, cur
		}
		prev = nextHead
		tail = nextHead
	}
	return node.Next
}

/**
* 给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
* 示例 1:
* 输入: 1->2->3->4->5->NULL, k = 2
* 输出: 4->5->1->2->3->NULL
* 解释:
* 向右旋转 1 步: 5->1->2->3->4->NULL
* 向右旋转 2 步: 4->5->1->2->3->NULL
* 示例 2:
* 输入: 0->1->2->NULL, k = 4
* 输出: 2->0->1->NULL
* 解释:
* 向右旋转 1 步: 2->0->1->NULL
* 向右旋转 2 步: 1->2->0->NULL
* 向右旋转 3 步: 0->1->2->NULL
* 向右旋转 4 步: 2->0->1->NULL
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	fast, slow, c := head, head, 0
	for fast != nil {
		fast = fast.Next
		c++
	}
	k = k % c
	if k == c || k == 0 {
		return head
	}
	fast, c = head, 0
	for c <= k {
		fast = fast.Next
		c++
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	fast.Next = head
	// 拆分循环链表
	res := slow.Next
	slow.Next = nil
	return res
}

func printList(head *ListNode) {
	node := head
	for node != nil && node.Next != head {
		println(node.Val)
		node = node.Next
	}
}
