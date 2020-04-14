package main

import "fmt"

func main() {
	str := "abcabecdd"
	fmt.Println(lengthOfLongestSubstring(str))
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
