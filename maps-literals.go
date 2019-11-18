package main

import "fmt"

type Location struct {
	lat float64
	long float64
}

var m=map[string]Location{
	"fb": Location{
		40.2,78.3,
	},
	"google": Location{
		99.33,32.23,
	},
}

func main(){
	fmt.Println(m)
}

