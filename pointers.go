package main

import "fmt"

func main(){
	i :=42
	j := 2701
	p := &i   // Point to i
	fmt.Println(*p) // Print i through the pointer
	*p = 21	  // Set i through the pointer
	fmt.Println(i)

	p = &j    // Point p to j
	fmt.Println(*p)	// Print j through the pointer p
	*p = *p / 37 	// Setting the value of j through pointer p
	fmt.Println(j)
	
}
