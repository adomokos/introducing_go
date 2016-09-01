package main

import "fmt"

/*
func main() {
  var x string = "Hello, World"
  fmt.Println(x)
}

func main() {
  var x string
  x = "first "
  fmt.Println(x)
  x = x + "second"
  fmt.Println(x)
}

func main() {
  x := "Hello, World"
  fmt.Println(x)
}

func main() {
  x := "Max"
  fmt.Println("My dog's name is ", x)
}

func main() {
  const x string = "Hello, World"
  fmt.Println(x)
}
*/

func main() {
  fmt.Print("Enter a number: ")
  var input float64
  fmt.Scanf("%f", &input)

  output := input * 2

  fmt.Println(output)
}
