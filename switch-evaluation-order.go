package main

import "fmt"
import "time"

func main(){
	fmt.Println("When's Saturday")

	today := time.Now().Weekday()
	fmt.Println("Today is :",today)
	switch time.Saturday{
		case today + 0:
			fmt.Print("Today")
		case today + 1:
			fmt.Print("Tomorrow")
		case today + 2:
			fmt.Print("In two days")
		default:
			fmt.Print("Too far away")
		}
}
