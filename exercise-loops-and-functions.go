package main

import (
	"fmt"
	"math"
)
func sqrt(num float64) float64{
	z := 1.0
	for i:=0;i<10;i++{
		z -= (z * z -num)/ (2*z)
	}
	return z
}

func main(){
	fmt.Println("sqrt returns : ", sqrt(79983748379))
	fmt.Println("math.Sqrt returns : ",math.Sqrt(79983748379))
}
