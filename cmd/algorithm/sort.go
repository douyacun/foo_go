package main

import "fmt"

func QSort(arr []int, start, end int) {
	if start < end {
		m, l, r := start, start+1, end
		for l < r {
			for arr[l] <= arr[m] && l < r {
				l++
			}
			for arr[r] > arr[m] && l < r {
				r--
			}
			arr[l], arr[r] = arr[r], arr[l]
		}
		if arr[l] > arr[m] {
			l--
		}
		arr[m], arr[l] = arr[l], arr[m]
		QSort(arr, start, l-1)
		QSort(arr, l+1, end)
	}
}

func main() {
	arr := []int{3, 5, 6, 1, 0, 9, 3, 3, 8, 7, 4}
	QSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
