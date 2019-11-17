package main

import "fmt"
import "math"

func main(){
	var x , y int = 3, 4
	var z float64 = math.Sqrt(float64(x *x + y*y))
	var a uint= uint(z)
	fmt.Println(x , y , z , a)
}
