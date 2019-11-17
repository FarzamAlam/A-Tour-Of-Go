package main

import "fmt"


func main(){
	f := fibonnaci()
	for i:= 0;i<10;i++{
		fmt.Println(f())
	}

}

func fibonnaci() func() int{
	first := 0
	second := 1

	return func()int {
		ret := first
		first, second = second , first+second
		return ret
	}
}
