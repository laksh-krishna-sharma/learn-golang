package main

import (
	"errors"
	"fmt"
)

func main() {
	printHello()
	printMe("laksh")
	fmt.Println(division(4, 0))
}

func printHello() {
	fmt.Println("Hello!")
}

func printMe(name string) {
	fmt.Println("Hello, " + name + "!")
}

func division(nr int32, dr int32) (int32, int32, error) {
	var err error
	if dr == 0 {
		err = errors.New("Cannot divide by 0")
		return nr, dr, err
	}
	res := nr / dr
	rem := nr % dr
	return res, rem, err
}
