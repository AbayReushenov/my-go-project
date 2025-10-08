package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y)) +.01
	var z uint = uint(f)
	fmt.Println(x, y, z, f)  // 3 4 5 5.01
}

// Type conversions
// The expression T(v) converts the value v to the type T.

// Some numeric conversions:

// var i int = 42
// var f float64 = float64(i)
// var u uint = uint(f)
// Or, put more simply:

// i := 42
// f := float64(i)
// u := uint(f)
