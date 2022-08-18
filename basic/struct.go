package main
import ("fmt")

type Circle struct {
    x float64
    y float64
    z float64
}
 
type Rectangle struct {
    x1 float64
    y1 float64
    x2 float64
    y2 float64
}

func main(){
	// var len = new(Circle)

	var test Circle
	fmt.Scan(&test.x)
	fmt.Scan(&test.y)
	fmt.Scan(&test.z)
	fmt.Println(test)
	fmt.Println("this is end of program")
	
}