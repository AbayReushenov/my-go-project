package main

import "fmt"

func main() {
    var i int
    j := i // j is an int
    fmt.Printf("i is of type %T\n", i)
    fmt.Printf("j is of type %T\n", j)

    ii := 42           // int
    f := 3.142        // float64
    g := 0.867 + 0.5i // complex128

    fmt.Printf("ii is of type %T\n", ii)
    fmt.Printf("f is of type %T\n", f)
    fmt.Printf("g is of type %T\n", g)

	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
}

// i is of type int
// j is of type int
// ii is of type int
// f is of type float64
// g is of type complex128
// v is of type int
