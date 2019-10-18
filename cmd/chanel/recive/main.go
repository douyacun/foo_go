package recive

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 5)
	quit := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 1; i < 30; i++ {
		wg.Add(1)
		go func(f int) {
			defer wg.Done()
			ch <- f * 2
		}(i)
	}
	go end(ch, quit)
	wg.Wait()
	close(ch)
}

func end(ch chan int, quit <-chan bool)  {
	for {
		x, ok := <-ch
		if !ok {
			return
		}
		fmt.Println(x)
	}
}
