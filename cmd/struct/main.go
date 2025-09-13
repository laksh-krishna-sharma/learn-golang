package main

import "fmt"

// Structs in Go are user-defined types that group related fields.
type Person struct {
	Name string // Field for person's name
	Age  int    // Field for person's age
}

// You can embed structs to achieve composition.
type Employee struct {
	Person     // Embedding Person struct
	EmployeeID int
}

// Interfaces define method sets. Any type that implements all methods is said to satisfy the interface.
type Greeter interface {
	Greet() string // Any type with Greet() string method implements Greeter
}

// Implementing the Greeter interface for Person
func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

func main() {
	// Creating a struct instance
	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p.Name, p.Age)

	// Using embedded struct
	e := Employee{Person: Person{Name: "Bob", Age: 40}, EmployeeID: 1234}
	fmt.Println(e.Name, e.Age, e.EmployeeID)

	// Using interface
	var g Greeter = p // Person implements Greeter
	fmt.Println(g.Greet())

	// Anonymous struct example
	anon := struct {
		Field1 int
		Field2 string
	}{Field1: 1, Field2: "test"}
	fmt.Println(anon)
}

// NOTES:
// - Structs are value types; assignment copies all fields.
// - Use pointers to structs for mutability and efficiency.
// - Interfaces are satisfied implicitly (no 'implements' keyword).
// - Methods can have pointer or value receivers.
// - Zero value of a struct has all fields set to their zero values.
