package main

import "fmt"

func main() {
  x := []int{48, 96, 86, 68,
            57, 82, 63, 70,
            37, 34, 83, 27,
            19, 97,  9, 17,
            }

  var y int = 0
  for i := 0; i < len(x); i++ {
    if y == 0 || x[i] < y {
      y = x[i]
    }
  }

  fmt.Println("Smallest", y)

  z := [6]string{"a", "b", "c", "d", "e", "f"}
  fmt.Println(z[2:5]) // slice includes low index value but not high index value

  xs := []float64{98,93,77,82,83}
  total := 0.0
  for _, v := range xs {
    total += v
  }
  fmt.Println(total/float64(len(xs)))
}
