package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}


// aaaaa@aaaaa-GF63-Thin-9SCXR:~/my-go-project$ go run main.go
// counting
// done
// 9
// 8
// 7
// 6
// 5
// 4
// 3
// 2
// 1
// 0

// Стекирование отложенных вызовов
// Отложенные вызовы функций помещаются в стек. При возврате из функции её отложенные вызовы выполняются в порядке «последним вошёл — первым вышел».
