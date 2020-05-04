package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%d", myAtoi("2147483648"))
}

// 最长子串
// cdd  输出 4 abcd
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var (
		res, i, j int
	)
	m := make(map[byte]struct{}, len(s))
	for i < len(s) && j < len(s) {
		if _, ok := m[s[j]]; ok {
			delete(m, s[i])
			i++
		} else {
			m[s[j]] = struct{}{}
			j++
		}
		if j-i > res {
			res = j - i
		}
	}
	return res
}

// 中心扩散法
// 最长回文子串
// 输入: "cbccd"
// 输出: "cbc"
// 1. 有中心
// 2. 中心在2个重复字符串中间
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	var start, end int
	for i := 0; i < len(s); i++ {
		l1 := expandAroundCenter(s, i, i)   // 有中心
		l2 := expandAroundCenter(s, i, i+1) // 无中心
		max := int(math.Max(float64(l1), float64(l2)))
		if max > end-start {
			start = i - (max-1)/2
			end = i + max/2
		}
	}
	return s[start : end+1]
}

// 输入: "cbabcd"
// 输出: 5
// l = 1 r = 3
// l = 0 r = 4
// 3 4
// 1 2 3 4
// i = 2 len = 3
// 3 - 1 / 2
func expandAroundCenter(s string, left, right int) int {
	l, r := left, right
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}

// "   -42" -42
func myAtoi(s string) int {
	var (
		n int8 = 0
		c = 0
	)
	for i := 0; i < len(s); i++ {
		if n == 0 {
			if s[i] == '-' {
				n = -1
			} else if s[i] == '+' {
				n = 1
			} else if s[i] >= '0' && s[i] <= '9' {
				n = 1
				c = c*10 + int(s[i]-'0')
			} else if s[i] == ' ' {
				continue
			} else {
				return 0
			}
		} else {
			if s[i] >= '0' && s[i] <= '9' {
				c = c*10 + int(s[i]-'0')
				if c >= 2<<30 {
					if n > 0 {
						return 2<<30 - 1
					} else {
						return 0 - 2<<30
					}
				}
			} else {
				break
			}
		}
	}
	return c * int(n)
}
