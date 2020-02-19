package main

import (
	"fmt"
	"math"
)

func _Sqrt(x float64) float64 {
	z := float64(1)
	for i:= 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)
		fmt.Printf("Iteration %v, value = %v\n", i, z)
	}
	return z
}

func Sqrt(x float64) float64 {
  z := float64(1)
  for {
    t := z
    z -= (z*z - x) / (2*z)
    fmt.Printf("z %v, t = %v\n", z, t)
    if math.Abs(t-z) < 1e-6 {
      break
    }
  }
  return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
