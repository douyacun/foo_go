package main

import "fmt"

func main() {
	a := []int{1, 3, 5, 5}
	b := a[2:len(a)]
	fmt.Println(b, cap(a))
}
