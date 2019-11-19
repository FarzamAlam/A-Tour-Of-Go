package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error()string{
	return fmt.Sprintf("cannot sqrt a negative number: %v",float64(e))
}

func sqrt(x float64)(float64,error){
	if x<0{
		n := ErrNegativeSqrt(x)
		return 0,n
	}
	return math.Sqrt(x),nil
}

func main(){
	sq, err := sqrt(2)
	fmt.Println(sq,err)
	
	_,err =sqrt(-2)
	fmt.Println(err)
}

