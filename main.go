package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

// aaaaa@aaaaa-GF63-Thin-9SCXR:~/my-go-project$ go run main.go
// Hello 世界
// Happy 3.14 Day
// Go rules? true
