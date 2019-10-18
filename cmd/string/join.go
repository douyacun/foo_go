package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"
)

func main() {
	a := make([]string, 100000)
	for i := 1; i <= 100000; i++ {
		a = append(a, strconv.Itoa(i))
	}
	fmt.Println(unsafe.Sizeof(a))
	wg := sync.WaitGroup{}
	wg.Add(3)
	go plus(&wg, a)
	go join(&wg, a)
	go buffer(&wg, a)
	wg.Wait()
}

func plus(wg *sync.WaitGroup, a []string) {
	defer wg.Done()
	start := time.Now()
	var str string
	for _, i := range a {
		str = str + i
	}
	latency := time.Now().Sub(start)
	fmt.Printf("plus latency: %s len: %d\n", latency, len(str))
}

func join(wg *sync.WaitGroup, a []string) {
	defer wg.Done()
	start := time.Now()
	str := strings.Join(a, "")
	latency := time.Now().Sub(start)
	fmt.Printf("join latency: %s len: %d\n", latency, len(str))
}

func buffer(wg *sync.WaitGroup, a []string) {
	defer wg.Done()
	start := time.Now()
	var str = bytes.NewBufferString("")
	for _, i := range a {
		str.WriteString(i)
	}
	latency := time.Now().Sub(start)
	fmt.Printf("buffer latency: %s len: %d\n", latency, len(str.String()))
}
