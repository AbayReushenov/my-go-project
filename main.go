package main

import "fmt"

func main() {
	sum := 0
	fmt.Println(sum)

	for i := range 10 {
			fmt.Println(i, sum)

		sum += i

		fmt.Println(i, sum)
	}

	fmt.Println(sum)
}

// aaaaa@aaaaa-GF63-Thin-9SCXR:~/my-go-project$ go run main.go
// 0
// 0 0
// 0 0
// 1 0
// 1 1
// 2 1
// 2 3
// 3 3
// 3 6
// 4 6
// 4 10
// 5 10
// 5 15
// 6 15
// 6 21
// 7 21
// 7 28
// 8 28
// 8 36
// 9 36
// 9 45
// 45
