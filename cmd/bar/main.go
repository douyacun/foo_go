package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	b := make([]int, len(a))
	copy(b, a)

	b = append(b[:0], b[1:]...)
	fmt.Println(a, b)
}