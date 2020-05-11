package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}

	c := arr[:0]

	fmt.Println(c, len(c), cap(c))
}
