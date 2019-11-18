package main

import "fmt"

func do (i interface{}){
	switch v:=i.(type){
	case int:
		fmt.Printf("Type = %T,Value =%v\n",v,v)
	case string:
		fmt.Printf("Type = %T,Value =%v\n",v,v)
	default:
		fmt.Printf("I don't know about this type:%T\n",v)
	}
}

func main(){
	do(12)
	do("hello")
	do(true)
}
