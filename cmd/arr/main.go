package main

import "fmt"

func main() {
	// ------------------- ARRAYS -------------------
	var intArr [3]int32
	intArr[1] = 123
	fmt.Println("Array element at index 0:", intArr[0])
	fmt.Println("Array slice [1:3]:", intArr[1:3])

	// array stores elements in contiguous memory
	intArr2 := [3]int32{1, 2, 3}
	fmt.Println("Addresses in intArr2:")
	fmt.Println(&intArr2[0], &intArr2[1], &intArr2[2])

	intArr3 := [...]int32{1, 2, 3, 4}
	fmt.Println("intArr3[0]:", intArr3[0])

	var intArr4 [3]int32 = [3]int32{5, 6, 7}
	fmt.Println("intArr4[2]:", intArr4[2])

	// ------------------- SLICES -------------------
	// Slice from array
	slice1 := intArr3[1:3] // elements at index 1 and 2
	fmt.Println("Slice1:", slice1)

	// Slice literal
	slice2 := []int{10, 20, 30}
	fmt.Println("Slice2:", slice2)

	// Append to slice
	slice2 = append(slice2, 40, 50)
	fmt.Println("Slice2 after append:", slice2)

	// ------------------- MAPS -------------------
	// Map declaration
	studentGrades := map[string]int{
		"Alice": 90,
		"Bob":   85,
	}
	fmt.Println("Student Grades:", studentGrades)

	// Add new entry
	studentGrades["Charlie"] = 95
	fmt.Println("Updated Student Grades:", studentGrades)

	// Access key safely
	grade, exists := studentGrades["David"]
	if exists {
		fmt.Println("David's grade:", grade)
	} else {
		fmt.Println("David not found")
	}

	// ------------------- LOOPS -------------------
	// For loop with index
	for i := 0; i < 5; i++ {
		fmt.Println("For loop i:", i)
	}

	// Range loop with array
	for idx, val := range intArr2 {
		fmt.Println("Index:", idx, "Value:", val)
	}

	// Range loop with map
	for name, grade := range studentGrades {
		fmt.Printf("%s scored %d\n", name, grade)
	}
}
