package main

import "fmt"

func main() {
	var num int64

	ch := make(chan int64, 3)

	for i := 0; i < 3; i++ {
		select {
		case ch <- 0:
		case ch <- 1:
		case ch <- 2:
		case ch <- 3:
		case ch <- 4:
		case ch <- 5:
		case ch <- 6:
		case ch <- 7:
		case ch <- 8:
		case ch <- 9:
		}
		num = num*10 + <-ch
	}

	fmt.Println(num)
}
