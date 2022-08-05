package main
import "fmt"

func fibonacci(num int) int{
	if num == 0 || num == 1{
		return 1;
	}
	return fibonacci(num-1) + fibonacci(num-2)
}

func main(){


	fmt.Println(fibonacci(10))
	fmt.Println("This is end of program")
}