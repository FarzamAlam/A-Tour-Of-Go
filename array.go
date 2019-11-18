package main

import "fmt"

func main(){
	var str [2]string

	str[0]="hello"
	str[1]="world!"

	fmt.Println(str[0],str[1])
	fmt.Println(str)

	primes := [5]int{1,2,3,7,11}
	fmt.Println(primes)
}
