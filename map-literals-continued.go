package main

import "fmt"

type Location struct{
	lat float64
	long float64
}
var m = map[string]Location{
	"fb":{343.5434,654.0},
	"google":{-54,433.5},
}

func main(){
	fmt.Println(m)
}
