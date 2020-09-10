package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().AddDate(0, -1, 0).Format("2006-01"))
}
