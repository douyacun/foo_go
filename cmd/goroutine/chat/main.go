package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

type client chan<- string

var (
	entering  = make(chan client)
	leaving   = make(chan client)
	messages  = make(chan string)
	clientMap = make(map[string]string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(listener)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				if !clients[cli] {
					continue
				}
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	addr := conn.RemoteAddr().String()
	input := bufio.NewScanner(conn)
	ch <- "给自己起一个名字和大家认识一下:"
	var name = ""
	if input.Scan() {
		name = input.Text()
	}
	clientMap[addr] = name
	ch <- fmt.Sprintf("welcome %s", name)
	entering <- ch
	messages <- fmt.Sprintf("系统: %s加入房间", name)

	abort := make(chan string)

	go func() {
		for {
			select {
			case <-time.After(1 * time.Minute):
				conn.Close()
			case str := <-abort:
				messages <- str
			}
		}
	}()

	for input.Scan() {
		str := input.Text()
		if str == "exit" {
			break
		}
		if len(str) > 0 {
			abort <- fmt.Sprintf("%s: %s", name, input.Text())
		}
	}
	leaving <- ch
	messages <- fmt.Sprintf("系统: %s离开了", name)
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}
