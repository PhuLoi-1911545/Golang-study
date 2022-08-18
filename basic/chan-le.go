package main
import "fmt"

func main(){
	for  i := 0; i < 11; i++{
		if i%2 == 0{
			fmt.Println("chan")
		} else{	
			fmt.Println("le")
		}
	}
}