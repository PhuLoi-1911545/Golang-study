package main
import "fmt"

func main(){
	var x *int
	var y = 321321
	x = &y
	y++
	// var z =new(int) 
	

	fmt.Println(*x)
}