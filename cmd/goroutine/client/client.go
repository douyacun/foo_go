package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Printf("tcp dial connect failed, %s", err)
	}
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, conn io.Reader) {
	_, err := io.Copy(dst, conn)
	if err != nil {
		fmt.Println(err)
		return
	}
}
