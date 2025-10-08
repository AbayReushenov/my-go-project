package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
// aaaaa@aaaaa-GF63-Thin-9SCXR:~/my-go-project$ go run main.go
// My favorite number is 4
// aaaaa@aaaaa-GF63-Thin-9SCXR:~/my-go-project$ go run main.go
// My favorite number is 6
// aaaaa@aaaaa-GF63-Thin-9SCXR:~/my-go-project$ go build main.go
// aaaaa@aaaaa-GF63-Thin-9SCXR:~/my-go-project$ ./main
// My favorite number is 9
// aaaaa@aaaaa-GF63-Thin-9SCXR:~/my-go-project$ ./main
// My favorite number is 7
// aaaaa@aaaaa-GF63-Thin-9SCXR:~/my-go-project$
