package main

import (
	"fmt"
)

func main() {

	fmt.Println(split(11))
}

func split(sum int) (x, y int) {
	x = sum / 2
	y = sum - x
	return
}
