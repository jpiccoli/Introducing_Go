package main

import "fmt"

func main() {
  square1 := func(p int) int {
    return p * p
  }

  fmt.Println("Square1 9 =", square1(9))
  fmt.Println("Square1 11 =", square1(11))

  y := factorial(7)
  fmt.Println("factorial =", y)

  cubeValue := makeCube(7)
  fmt.Println("Cube =", cubeValue())

  fmt.Println("---------------------------")

  // Item 1
  slice := []int{27, 29, 31}
  sum1 := sum(slice)
  fmt.Println("Item 1. sum =", sum1)

  // Item 2
  arg1, arg2 := half(16)
  fmt.Println("Item 2a. val = 16, arg1 =", arg1, "arg2 =", arg2)
  arg1, arg2 = half(17)
  fmt.Println("Item 2b. val = 17, arg1 =", arg1, "arg2 =", arg2)

  // Item 3
  m := max(90,91,92,93,99,94,95,96,97,98)
  fmt.Println("Item 3. m =", m)

  //Item 5
  f := fibonacci(12)
  fmt.Println("Item 5. Fibonacci =", f)

  // Item 9
  xPtr := new(int)
  *xPtr = 29
  fmt.Println("Item 9. xPtr =", *xPtr)

  // Item 10
  f1 := 1.5
  square(&f1)
  fmt.Println("Item 10. Square", f1)

  // Item 11
  a := int(1)
  b := int(2)
  swap(&a, &b)
  fmt.Println("Item 11. a =",a,",b =", b)
}

// Example functions
func factorial(x int) int {
  if x == 0 {
    return 1
  }
  return  x * factorial(x-1)
}

func makeCube(q int) func() int {
  r := int(q)
  return func() (ret int) {
    ret = r * r * r
    return
  }
}

// Item 1
func sum(x []int) int {
  total := int(0)
  for _, j := range x {
    total += j
  }
  return total
}

// Item 2
func half(x int) (int, bool) {
  if x % 2 == 0 {
    return 1, true
  } else {
    return 0, false
  }
}

// Item 3
func max(args...int) int {
  max_value := int(0)
  for _, v := range args {
    if max_value == 0 || max_value < v {
      max_value = v
    }
  }
  return max_value
}

// Item 5
func fibonacci (v int) int {
  if v == 0 {
    return 0
  } else if v == 1 {
    return 1
  } else {
    return fibonacci(v-1) + fibonacci(v-2)
  }
}

// Item 10
func square(x *float64) {
  *x = *x * *x
}

// Item 11
func swap(x *int, y *int) {
    temp := *x
    *x = *y
    *y = temp
}
