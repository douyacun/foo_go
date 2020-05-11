package main

import "fmt"

func QSort(arr []int, start, end int) {
	if start < end {
		mid := arr[start]
		var (
			i = start + 1
			j = end
		)
		for i <= j {
			for arr[i] <= mid && i < j {
				i++
			}
			arr[i], arr[j] = arr[j], arr[i]
			j--
		}
		if arr[i] > mid {
			i--
		}
		arr[start], arr[i] = arr[i], arr[start]
		QSort(arr, start, i-1)
		QSort(arr, i+1, end)
	}
}

func main() {
	arr := []int{3, 5, 6, 1, 0, 9, 3, 3, 8, 7, 4}
	QSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
