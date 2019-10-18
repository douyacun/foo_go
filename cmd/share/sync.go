package share

import "fmt"

var x, y int

func syncVariable() {
	go func() {
		x = 1
		fmt.Print("y: ", y, " ")
	}()
	go func() {
		y = 1
		fmt.Print("x: ", x, " ")
	}()
}
