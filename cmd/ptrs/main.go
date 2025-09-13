package main

import "fmt"

// Pointers in Go store the memory address of a value.
// The zero value of a pointer is nil.

func main() {
	// Declaring a pointer to an int
	var p *int
	fmt.Println(p) // nil (no address assigned)

	// Getting the address of a variable using &
	x := 10
	p = &x          // p now points to x
	fmt.Println(p)  // prints address
	fmt.Println(*p) // dereferencing: prints 10

	// Changing value via pointer
	*p = 20
	fmt.Println(x) // x is now 20

	// Passing pointer to a function allows modification of original value
	increment(&x)
	fmt.Println(x) // x is now 21

	// Pointers with structs
	type Person struct {
		Name string
		Age  int
	}
	person := Person{"Alice", 30}
	personPtr := &person
	personPtr.Age = 31 // Can use . to access fields via pointer
	fmt.Println(person)
}

func increment(n *int) {
	*n++ // increment value at address n
}

// NOTES:
// - Use & to get address, * to dereference.
// - Pointers are useful for large structs, slices, maps, and for mutating values in functions.
// - Go has no pointer arithmetic (unlike C/C++).
// - Slices, maps, and channels are reference types (internally use pointers).
