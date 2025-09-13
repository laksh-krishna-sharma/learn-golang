package main

import (
	"fmt"
	"strings"
)

func main() {
	// Strings in Go are immutable sequences of bytes (UTF-8 encoded)
	var myString = "Hello, Go!" // Declare a string
	var myIndex = myString[0]   // Accessing the first byte (not rune/character)

	// Print the value and type of myIndex (byte value of 'H')
	fmt.Printf("%v, %T\n", myIndex, myIndex)

	// Iterating over a string with range gives index (in bytes) and rune value (Unicode code point)
	for i, v := range myString {
		fmt.Println(i, v) // i: byte index, v: rune (int32)
	}

	// Convert string to a slice of runes to handle Unicode characters properly
	var myRune = []rune("Hello, Go!")

	// Iterating over rune slice gives index and Unicode code point
	for i, v := range myRune {
		fmt.Println(i, v) // i: rune index, v: rune value
	}

	// String concatenation using slices and strings.Builder
	var strSlice = []string{"H", "e", "l", "l", "o"} // Slice of strings
	var catSlice = ""                                // Concatenated string (inefficient for many items)
	var stringBuilder strings.Builder                // Efficient way to build strings

	for i := range strSlice {
		catSlice += strSlice[i]                // Concatenate using += (creates new string each time)
		stringBuilder.WriteString(strSlice[i]) // Append to Builder (more efficient)
	}

	var catSlice2 = stringBuilder.String() // Get the final string from Builder

	fmt.Println(catSlice2) // Output: Hello
	fmt.Println(catSlice)  // Output: Hello

	// Converting a byte slice to string
	var byteSlice = []byte{'H', 'e', 'l', 'l', 'o'} // Slice of bytes (ASCII values)
	var catByte = string(byteSlice)                 // Convert byte slice to string

	fmt.Println(catByte) // Output: Hello
}
