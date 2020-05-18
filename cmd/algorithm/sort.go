package main

import "fmt"

func QSort(nums []int, start, end int) {
	if start < end {
		mid := start
		left := start + 1
		right := end
		for left < right {
			for left < right && nums[left] <= nums[mid] {
				left++
			}
			for left < right && nums[right] > nums[mid] {
				right--
			}
			nums[left], nums[right] = nums[right], nums[left]
			right--
		}
		if nums[left] > nums[mid] {
			left--
		}
		nums[left], nums[mid] = nums[mid], nums[left]
		QSort(nums, start, left - 1)
		QSort(nums, left + 1, end)
	}
}

func main() {
	arr := []int{3, 5, 6, 1, 0, 9, 3, 3, 8, 7, 4}
	QSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
