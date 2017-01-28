package main

import "fmt"

func main() {
  for i := 1; i < 10; i++ {
    // fmt.Println(i)
    if i % 2 == 0 {
      fmt.Println(i, "even")
    } else {
      fmt.Println(i, "odd")
    }

    switch i {
    case 1: fmt.Println("One")
    case 2: fmt.Println("Two")
    default: fmt.Println("Greater than Two")
    }
  }

  for i := 1; i <= 100; i++ {
    if i % 3 == 0 {
      fmt.Print(i, ",")
    }
  }

  fmt.Println("")

  for i := 1; i <= 100; i++ {
    if i % 3 == 0 && i % 5 == 0 {
      fmt.Println("FizzBuzz")
    } else if i % 3 == 0 {
      fmt.Println("Fizz")
    } else if i % 5 == 0 {
      fmt.Println("Buzz")
    } else {
      fmt.Println(i)
    }
  }
}
