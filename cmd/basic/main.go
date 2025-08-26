package main
import "fmt"
import "unicode/utf8"


func foo() string {
    return "Hello from foo!"
}

func main(){

  fmt.Println("Hello, World!")

  var intNum int = 37456
  fmt.Println(intNum)

  var floatNum float64 = 12345.98
  fmt.Println(floatNum)

  var intNum32 int32 = 1234
  var floatNum32 float32 = 1234.56
  var result float32 = floatNum32 + float32(intNum32)
  fmt.Println(result)

  var intNum1 int32 = 3
  var intNum2 int32 = 2

  fmt.Println(intNum1/intNum2)
  fmt.Println(intNum1%intNum2)

  var myString string = "Hello," + " " + "World!"
  fmt.Println(utf8.RuneCountInString(myString)) // return length of string

  fmt.Println(len("test")) // return byet size

  var myRune rune = 'q'
  fmt.Println(myRune)

  var myBoolean bool = true
  fmt.Println(myBoolean)

  // to drop var keyword use :=

  myName := "Laksh Krishna Sharma"
  fmt.Println(myName)

  var1, var2 := 1, 2
  fmt.Println(var1, var2)

  var myFunc string = foo()
  fmt.Println(myFunc)

  const pi float32 = 3.1415
  fmt.Println(pi)
}
