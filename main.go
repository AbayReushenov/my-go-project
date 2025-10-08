package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
// aaaaa@aaaaa-GF63-Thin-9SCXR:~/my-go-project$ go run main.go
// 55
