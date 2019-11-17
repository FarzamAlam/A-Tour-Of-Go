package main

import "fmt"
import "runtime"

func main(){
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os{
		case "darwin":
			fmt.Print("darwin")
		case "linux":
			fmt.Print("linux")
		case "windows":
			fmt.Print("windows")
		default :
			fmt.Print("Nothing!")
	}
}
