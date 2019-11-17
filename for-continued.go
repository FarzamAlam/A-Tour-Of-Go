package main

import "fmt"

func main(){
	var sum int = 0
	var i int = 0
	for ;sum < 1000; {
		sum +=10
		i +=1
	}

	fmt.Println("Iteration took: ",i)
}
