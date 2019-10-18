package spinner

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	n := flip(45)
	fmt.Println(n)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func flip(x int) int {
	if x < 2 {
		return x
	}
	return flip(x-1) + flip(x-2)
}
