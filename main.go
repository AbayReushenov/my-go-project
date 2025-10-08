package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())
}

// aaaaa@aaaaa-GF63-Thin-9SCXR:~/my-go-project$ ./main
// Hello, World!
// aaaaa@aaaaa-GF63-Thin-9SCXR:~/my-go-project$ go build main.go
// aaaaa@aaaaa-GF63-Thin-9SCXR:~/my-go-project$ ./main
// Welcome to the playground!
// The time is 2025-10-08 10:19:38.253886962 +0300 MSK m=+0.000220249
// aaaaa@aaaaa-GF63-Thin-9SCXR:~/my-go-project$ go run main.go
// Welcome to the playground!
// The time is 2025-10-08 10:20:19.063815585 +0300 MSK m=+0.000106998
// aaaaa@aaaaa-GF63-Thin-9SCXR:~/my-go-project$
