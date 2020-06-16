package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

/**
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的数字可以无限制重复被选取。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum

回溯+剪枝
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

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return combinationSum3(candidates, target, 0)
}
/**
申请新的slice作为子集的candidates需要占用额外的内存，可以使用index来节省
 */
func combinationSum3(candidates []int, target int, index int) [][]int {
	res := make([][]int, 0)
	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}
		reminder := target - candidates[i]
		if reminder == 0 {
			res = append(res, []int{candidates[i]})
		} else if reminder > 0 {
			expect := combinationSum3(candidates, reminder, i+1)
			for _, j := range expect {
				m := []int{candidates[i]}
				m = append(m, j...)
				res = append(res, m)
			}
		}
	}

	return res
}
