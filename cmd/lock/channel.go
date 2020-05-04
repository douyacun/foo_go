package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sema    = make(chan struct{})
	mu      = sync.Mutex{}
	balance = 200
)

func main() {
	go Deposit(1)
	go Deposit(-1)

	fmt.Println(cap(sema))

	time.Sleep(500 * time.Millisecond)
	fmt.Println(balance)
}

func Deposit(amout int) {
	mu.Lock()
	defer mu.Unlock()
	a := balance
	time.Sleep(10 * time.Millisecond)
	balance = a + amout
}

func Balance() int {
	return balance
}
