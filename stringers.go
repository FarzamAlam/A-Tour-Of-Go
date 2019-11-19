package main

import "fmt"

type Person struct {
	Name string
	Age int
}
func (p Person) String()string{
	return fmt.Sprintf("%v is %v years old",p.Name,p.Age)
}

func main(){
	p1 := Person{"Anil",38}
	p2 := Person{"Aran",19}

	fmt.Println(p1,p2)
}
