package main
import "fmt"

func main(){
	arr := make([]int,10)
	for i := range(arr) {
		fmt.Println(i)
		arr[i] = i
	}

	fmt.Println("New kind of array")

	arr2 := arr[2:5]
	
	for i := range(arr2) {
		fmt.Println(arr2[i])
	}
}