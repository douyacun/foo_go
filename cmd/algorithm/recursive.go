package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{0, 0, 0}, 0))
}

/**
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的数字可以无限制重复被选取。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum
 */
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	for k, v := range candidates {
		reminder := target - v
		if reminder == 0 {
			res = append(res, []int{v})
		} else if reminder > 0 {
			expect := combinationSum(candidates[k:len(candidates)], reminder)
			for _, j := range expect {
				t := []int{v}
				t = append(t, j...)
				res = append(res, t)
			}
		}
	}
	return res
}
