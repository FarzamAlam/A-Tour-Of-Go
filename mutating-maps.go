package main

import "fmt"

func main(){
	m := make(map[string]int)
	m["a"] = 10
	fmt.Println(m["a"])
	
	m["a"] = 99
	fmt.Println(m["a"])
	
	_, ok := m["a"]
	fmt.Println(ok)
	
	delete(m,"a")
	_,ok =m["a"]
	fmt.Println(ok)
}

