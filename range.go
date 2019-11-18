package main

import "fmt"

var list = []int{2,4,6,8,10}

func main(){
	for i,v := range list{
		fmt.Println(i,v)
	}
}

