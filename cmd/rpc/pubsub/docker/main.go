package main

import (
	"fmt"
	"github.com/docker/docker/pkg/pubsub"
	"time"
)

func main() {
	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	c := p.Subscribe()

	p.Publish("hi")

	msg := <-c
	fmt.Println(msg)
}
