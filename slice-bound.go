package main

import "fmt"

func main(){
	arr := [10]int{1,2,3,4,5,6,7,8,9,10}

	fmt.Println(arr[:])
	fmt.Println(arr[0:])
	fmt.Println(arr[0:10])
}
