package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(10, 15))
}

func add(x int, y int) int {
	return x + y
}
