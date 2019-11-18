package main

import (
	"fmt"
	"math"
)
type I interface {
	M()
}
type T struct {
     S string
}

type F float64

func (t *T) M(){
	fmt.Println(t.S)
}

func (f F) M(){
	fmt.Println(f)
}

func main(){
	var i I
	f := F(math.Pi)
	i = f
	i.M()
	describe(i)

	t := &T{"Hello"}
	i= t
	i.M()
	describe(i)
}
func describe(i I){
	fmt.Printf("(%v,%T)\n",i,i)
}
