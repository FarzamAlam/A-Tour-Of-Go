package main

import "fmt"

type Vertex struct {
	lat float64
	long float64
}

var m map[string]Vertex

func main(){
	m = make(map[string]Vertex)
	m["a"] = Vertex{1.0,2.0}
	m["b"] = Vertex{3,4.0} 
	fmt.Println(m["a"])
	fmt.Println(m["b"])
	fmt.Println(m)
}
