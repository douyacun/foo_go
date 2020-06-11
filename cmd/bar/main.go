package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Weight int
}

func main() {
	bar := person{Weight: 80}
	data, _ := json.Marshal(bar)
	fmt.Println(string(data))
}
