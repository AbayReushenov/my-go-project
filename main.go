package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
    fmt.Println(swap("Привет", "Мир"))
}
// aaaaa@aaaaa-GF63-Thin-9SCXR:~/my-go-project$ go run main.go
// world hello
