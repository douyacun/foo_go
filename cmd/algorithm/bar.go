package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// 字符数统计
func letterCount(s string, code byte) int {
	var (
		letters = []byte(s)
		count   int
	)
	for _, v := range letters {
		if v == code {
			count++
		}
	}
	return count
}

func letterUnique(s string) string {
	var (
		letters   = []byte(s)
		container = make(map[byte]int)
		buf       = bytes.NewBuffer(nil)
	)
	for _, v := range letters {
		if _, ok := container[v]; !ok {
			buf.WriteByte(v)
		}
		container[v]++
	}
	return buf.String()
}

func letterCompress(s string) string {
	var (
		letters = []byte(s)
		buf     = bytes.NewBuffer(nil)
		length  = len(letters)
		c       = 1
	)
	for i, j := 0, 0; i < length; {
		j++
		for j < length && letters[i] == letters[j] {
			c++
			j++
		}
		if c > 1 {
			buf.WriteString(strconv.Itoa(c))
			buf.WriteByte(letters[i])
		} else {
			buf.WriteByte(letters[i])
		}
		i = j
		c = 1
	}
	return buf.String()
}

func arithmetic(s string) string {
	op := strings.Split(s, " ")
	if len(op) > 3 {
		return "0"
	}
	l, err := strconv.Atoi(op[0])
	if err != nil {
		return "0"
	}
	r, err := strconv.Atoi(op[2])
	if err != nil {
		return "0"
	}
	switch op[1] {
	case "+":
		return strconv.Itoa(l + r)
	case "-":
		return strconv.Itoa(l + r)
	case "/":
		return strconv.Itoa(l + r)
	case "*":
		return strconv.Itoa(l + r)
	}
	return "0"
}

func main() {
	var (
		s string
	)
	s = "https://www.douyacun.com"
	fmt.Println(letterCompress(s))
}
