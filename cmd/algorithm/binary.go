package main

import "fmt"

// 二分法 解题
func main() {
	nums := []int{2, 2}
	fmt.Println(searchRange(nums, 2))
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}
		if left+1 == right {
			break
		}
		if nums[left] <= nums[mid] {
			if nums[left] < target && nums[mid] > target {
				right = mid
			} else {
				left = mid
			}
		} else if nums[mid] < nums[right] {
			if nums[mid] < target && nums[right] > target {
				left = mid
			} else {
				right = mid
			}
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	index := -1
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			index = mid
			break
		}
		if nums[left] == target {
			index = left
			break
		}
		if nums[right] == target {
			index = right
			break
		}
		if left+1 == right {
			break
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if index == -1 {
		return []int{-1, -1}
	} else {
		left, right = index, index
		for left > 0 {
			if nums[left-1] == target {
				left = left - 1
			} else {
				break
			}
		}
		for right < len(nums)-1 {
			if nums[right+1] == target {
				right = right + 1
			} else {
				break
			}
		}
		return []int{left, right}
	}
}
