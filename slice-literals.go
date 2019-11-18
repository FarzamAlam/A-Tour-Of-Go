package main

import "fmt"

func main(){
	q := []int {1,3,5,7,11}
	fmt.Println(q)

	b := []bool {true,false,true,false}
	fmt.Println(b)

	s := []struct{
		i int
		c bool
	}{
	{1,true},
	{2,true},
	{3,false},
	{4,true},
	}

	fmt.Println(s)
}
