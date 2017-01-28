package main

import "fmt"

var w string = "String0"

func main() {
  var x string = "String1"
  var y string = "String2"
  fmt.Println(x == y)
  z := "String3"
  fmt.Println(z) // type inference
  f()
  fmt.Print("Enter a number: ")
  var input float64
  fmt.Scanf("%f", &input)
  output := input * 2
  fmt.Println(output)

  fmt.Print("Enter a temperature: ")
  var fahrenheit float64
  fmt.Scanf("%f", &fahrenheit)
  //fmt.Println("Fahrenheit =", fahrenheit)
  celsius := (fahrenheit - 32.0) * (5.0/9.0)
  fmt.Println(celsius)
}

func f() {
  fmt.Println(w)
}
