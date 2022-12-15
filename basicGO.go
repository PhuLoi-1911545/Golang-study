package main

import (
	"fmt"

)

//hello world code

// func main() {
// 	fmt.Println("Hello World")
// }

// variable
/* Các biến phải cùng type, đổi type phải cùng bit */

// func main() {
// 	// var numA int64 = 16;
// 	// var numA2 float32 = 19.9;
// 	// numB := 18.8;
// 	// numB2 := 3;

// 	var (
// 		numA int64 = 16
// 		numA2 float32 = 19.9
// 		numB float64 = 18.9
// 		numB2 int = 4
// 	)

// 	// var numA, numA2 =16, "abc"
// 	// numB, numB2 := "asdasd", 9

// 	// Const
// 	// const PI = 3.14
// 	// const A int = 1

// 	fmt.Println("Result is ", float64(numA) + numB) // int -> float, int64 -> float64
// 	fmt.Println("Result is ", numA + int64(numB))
// 	fmt.Print("Result is ", int32(numA2) + int32(numB2))
// }

// Formating output
// func main() {
// 	var i = 15.5
// 	// var i2 = 9
// 	var txt = "Hello World!"

// 	fmt.Printf("%v\n", i) 	//	Prints the value in the default format
// 	fmt.Printf("%#v\n", i) 	//	Prints the value in Go-syntax format
// 	fmt.Printf("%v%%\n", i)	//	Prints the % sign
// 	fmt.Printf("%T\n", i)	//	Prints the type of the value

// 	fmt.Printf("\n%v\n", txt)
// 	fmt.Printf("%#v\n", txt)
// 	fmt.Printf("%T\n", txt)

// 	fmt.Println("\n next type \n")
// 	var i2 = 13
// 	fmt.Printf("%b\n", i2)	//binary
// 	fmt.Printf("%d\n", i2)	// decimal
// 	fmt.Printf("%+d\n", i2)	//signed decimal
// 	fmt.Printf("%o\n", i2)	// hex
// 	fmt.Printf("%O\n", i2)	// hex with leading 0o
// 	fmt.Printf("%x\n", i2)	// Base 16, lowercase
// 	fmt.Printf("%X\n", i2)	// 	Base 16, uppercase
// 	fmt.Printf("%#x\n", i2)	// 	Base 16, with leading 0x
// 	fmt.Printf("%4d\n", i2)	// 	Pad with spaces (width 4, right justified)
// 	fmt.Printf("%-4d\n", i2)	// 	Pad with spaces (width 4, left justified)
// 	fmt.Printf("%04d\n", i2)	//	Pad with zeroes (width 4
// }

// array
// func main() {
// 	var arr1 = [3]string{"aa11","aa22","a222"}
// 	arr2 := [3]int{1,2,3}

// 	arr1[3] = "test2" //this will get error

// 	fmt.Println(arr1)
// 	fmt.Println(arr2);
// }

// slice
// func main() {
// 	arr := [1]int{1}
// 	sli := []int{1}
// 	fmt.Printf("size of array: %d byte and slice: %d byte \n", unsafe.Sizeof(arr), unsafe.Sizeof(sli))

// 	arr2 := [5]int{1,2,3,4,5}
// 	sli2 := arr2[:]
// 	fmt.Println(sli2)
// }



// Copy
// func main() {
// 	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
// 	// Original slice
// 	fmt.Printf("numbers = %v\n", numbers)
// 	fmt.Printf("length = %d\n", len(numbers))
// 	fmt.Printf("capacity = %d\n", cap(numbers))
	  
// 	// Create copy with only needed numbers
// 	neededNumbers := numbers[:len(numbers)-10]
// 	numbersCopy := make([]int, len(neededNumbers))
// 	fmt.Printf("type of numbersCopy: %T\n", numbersCopy)
// 	copy(numbersCopy, neededNumbers)  
  
// 	fmt.Printf("numbersCopy = %v\n", numbersCopy)
// 	fmt.Printf("length = %d\n", len(numbersCopy))
// 	fmt.Printf("capacity = %d\n", cap(numbersCopy))

//   }




// switch
func main() {
	// day := 5
	for day := 0; day <= 10; day++ 	{
		switch day {
			case 1,3,5:
			fmt.Println("Odd weekday")
			case 2,4:
			fmt.Println("Even weekday")
			case 6,7:
			fmt.Println("Weekend")
		default:
			fmt.Println("Invalid day of day number")
	   }
	}
}