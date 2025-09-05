package main

import "fmt"

func main() {
	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	// array stores elements in contigous manner
	intArr2 := [3]int32{1, 2, 3}
	fmt.Println(&intArr2[0])
	fmt.Println(&intArr2[1])
	fmt.Println(&intArr2[2])

	intArr3 := [...]int32{1, 2, 3, 4}
	fmt.Println(intArr3[0])

	var intArr4 [3]int32 = [3]int32{5, 6, 7}
	fmt.Println(intArr4[2])
}
