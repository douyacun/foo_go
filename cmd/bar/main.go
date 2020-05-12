package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func worker(ctx context.Context, id int, tasks chan int) {
	for {
		select {
		case t, ok := <-tasks:
			if ok {
				fmt.Printf("worker %d started task %d\n", id, t)
				time.Sleep(time.Second * 2)
				fmt.Printf("worker %d finished task %d\n", id, t)
			}
		case <-ctx.Done():
			fmt.Printf("worker %d is cancel\n", id)
			return
		default:
			fmt.Printf("worker %d is waiting for a task\n", id)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	tasks := make(chan int) // 队列
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	for wid := 1; wid <= runtime.NumCPU(); wid++ {
		go worker(ctx, wid, tasks)
	}
	for t := 1; t <= 10; t++ {
		tasks <- t
	}

	time.Sleep(10 * time.Second)
	close(tasks)
	cancel()
}
