package main
import "fmt"
import "time"

func main(){
	t := time.Now()
	switch {
		case t.Hour()<12:
			fmt.Println("Good Night")
		case t.Hour()<17:
			fmt.Println("Good Afternoon")
		default:
			fmt.Println("Good Night")
	}
}
