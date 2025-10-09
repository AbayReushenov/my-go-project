package main

import "fmt"

func main() {
	sum := 1
	for sum < 1030 { // As WHILE
		sum += sum
	}
	fmt.Println(sum) // 2048
}
