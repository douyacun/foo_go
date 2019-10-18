package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	p := NewPublisher(100*time.Millisecond, 10)
	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})
	p.Publish("Hello world")
	p.Publish("Hello golang")

	go func() {
		for msg := range all {
			fmt.Printf("all: %s\n", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Printf("golang: %s\n", msg)
		}
	}()

	time.Sleep(3 * time.Second)
}
