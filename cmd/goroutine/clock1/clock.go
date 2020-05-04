package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	clock2("localhost:8000")
}

func clock2(addr string) {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Printf("net listen failed, %s", err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("listener accept failed, %s\n", err)
			return
		}
		fmt.Printf("new tcp connect, %s\n", conn.RemoteAddr())
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	wg := sync.WaitGroup{}
	input := bufio.NewScanner(conn)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for input.Scan() {
			wg.Add(1)
			go echo(conn, input.Text(), 1*time.Second, wg)
		}
	}(&wg)
	wg.Wait()
	fmt.Printf("conn: %s quit\n", conn.RemoteAddr())
	conn.Close()
}

func echo(conn net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := fmt.Fprintln(conn, "\t", strings.ToUpper(shout))
	if err != nil {
		fmt.Printf("shout to conn failed, %s", shout)
		return
	}
	time.Sleep(delay)
	_, err = fmt.Fprintln(conn, "\t", shout)
	if err != nil {
		fmt.Printf("shout to conn failed, %s", shout)
		return
	}
	_, err = fmt.Fprintln(conn, "\t", strings.ToLower(shout))
	if err != nil {
		fmt.Printf("shout to conn failed, %s", shout)
		return
	}
}
