package main

import "fmt"

func main() {
	//nums := []int{3, 2, 4}
	//fmt.Println(twoSum(nums, 6))
	//arr1 := []int{0, 1, 2, 4}
	//arr2 := []int{3, 5, 7}

	//fmt.Println(arrayCombine(arr1, arr2, 2))
	//fmt.Println(findSubstring("barfoothefoobarman", []string{"bar", "foo"}))
	//fmt.Println(validParentheses("()"))
	fmt.Println(longestValidParentheses("()(()))"))
}

/**
 * 输入一个整数，输出该数二进制表示中 1 的个数。例如，把 9 表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。
 */
func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		if num%2 == 1 {
			res++
		}
		num = num / 2
	}
	return res
}

func intToBin(num uint32) uint64 {
	if num <= 1 {
		return uint64(num)
	}
	var (
		res   uint64
		times uint64 = 1
	)
	for num > 0 {
		res += uint64(num%2) * times
		num /= 2
		times *= 10
	}
	return res
}

/**
 * 两个有序数组从小到大排列，取两个数组合并后第K大的元素，要求O(1)空间复杂度
 * 如 a = {0, 1, 2, 4} b = {3, 5, 7} 第4大元素，返回3
 */
func arrayCombine(a, b []int, k int) int {
	if len(a) > len(b) {
		return arrayCombine(b, a, k)
	}

	if len(a) == 0 {
		return b[k-1]
	}

	if k == 1 {
		if a[0] < b[0] {
			return a[0]
		} else {
			return b[0]
		}
	}

	x, y := len(a), len(b)
	if len(a) > k/2 {
		x = k / 2
	}
	if len(b) > k/2 {
		y = k / 2
	}

	if a[x-1] < b[y-1] {
		return arrayCombine(a[x:], b, k-x)
	} else if a[x-1] > b[y-1] {
		return arrayCombine(a, b[y:], k-y)
	} else {
		return a[x-1]
	}
}

func findSubstring(s string, words []string) []int {
	res := make([]int, 0)
	if len(words) == 0 {
		return res
	}
	windowLength := len(words[0]) * len(words)
	for j := 0; j < len(s) && (len(s)-j) >= windowLength; {
		if includeSubStr(s[j:j+windowLength], words) {
			res = append(res, j)
		}
		j++
	}
	return res
}

func delSlice(t []string, i int) []string {
	if i == len(t) {
		return t[:i]
	} else {
		return append(t[:i], t[i+1:]...)
	}
}

func copySlice(src []string) []string {
	dst := make([]string, len(src))
	copy(dst, src)
	return dst
}

func includeSubStr(s string, words []string) bool {
	t := copySlice(words)
	i := 0
	for len(t) > 0 && i < len(s) {
		prev := len(t)
		for k, word := range t {
			if s[i:i+len(word)] == word {
				i += len(word)
				t = delSlice(t, k)
				break
			}
		}
		if len(t) == prev {
			return false
		}
	}
	if len(t) == 0 {
		return true
	}
	return false
}

func longestValidParentheses(s string) int {
	max := 0
	for i := 0; i < len(s)-1; {
		for j := i; j < len(s)-1; {
			c := j
			for c < len(s) - 1 && s[c+1] == s[j] {
				c++
			}
			end := 2*(c-j) + j + 1
			if end < len(s) && validParentheses(s[j:end+1]) {
				j = end + 1
				if j - i > max {
					max = j - i
				}
				if j >= len(s)-1 {
					i = j
				}
			} else {
				if i == j {
					i++
				} else {
					i = j
				}
				break
			}
		}
	}
	return max
}

func validParentheses(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}
