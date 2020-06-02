package main

import "fmt"

func main() {
	str := "()(()"
	fmt.Println(longestValidParentheses2(str))
}

func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}
	max := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i > 1 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else {
				if i-dp[i-1]-1 >= 0 {
					if s[i-dp[i-1]-1] == '(' {
						if i-dp[i-1]-1 >= 2 {
							dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
						} else {
							dp[i] = dp[i-1] + 2
						}
					}
				}
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

type stakeItem struct {
	val  interface{}
	next *stakeItem
}

type Stake struct {
	header *stakeItem
}

func NewStake() *Stake {
	return &Stake{}
}

func (s *Stake) Pop() interface{} {
	item := s.header
	s.header = s.header.next
	item.next = nil
	return item.val
}

func (s *Stake) Push(v interface{}) {
	item := &stakeItem{
		val:  v,
		next: s.header,
	}
	s.header = item
}

func (s *Stake) Empty() bool {
	return s.header == nil
}

func longestValidParentheses2(s string) int {
	if len(s) < 2 {
		return 0
	}
	max := 0
	d := NewStake()
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			d.Push(i)
		} else {
			if !d.Empty() {
				p := d.Pop().(int)
				if max > i - p + 1 {
					max = i - p + 1
				}
			}
		}
	}
	return max
}
