package main

import "fmt"

func main() {
    // Способ 1: var + тип
    var name string = "Golang"
    var age int = 15

    // Способ 2: var с автоопределением типа
    var language = "Go"
    var version = 1.21

    // Способ 3: короткое объявление (:=)
    developer := "Google"
    year := 2009

    // Несколько переменных одновременно
    var x, y int = 10, 20
    a, b := "hello", "world"

    // Константы
    const pi = 3.14159
    const greeting = "Привет"

    fmt.Println(name, age, language, version)
    fmt.Println(developer, year)
    fmt.Println(x, y, a, b)
    fmt.Println(pi, greeting)
}

// aaaaa@aaaaa-GF63-Thin-9SCXR:~/my-go-project$ go run main.go
// Golang 15 Go 1.21
// Google 2009
// 10 20 hello world
// 3.14159 Привет
// aaaaa@aaaaa-GF63-Thin-9SCXR:~/my-go-project$ go build main.go
// aaaaa@aaaaa-GF63-Thin-9SCXR:~/my-go-project$ ./main
// Golang 15 Go 1.21
// Google 2009
// 10 20 hello world
// 3.14159 Привет
