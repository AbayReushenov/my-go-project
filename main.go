package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

// hello
// world

//
// Defer
// A defer statement defers the execution of a function until the surrounding function returns.

// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
// Defer
// Оператор defer откладывает выполнение функции до тех пор, пока окружающая функция не вернёт управление.

// Аргументы отложенного вызова вычисляются немедленно, но вызов функции не выполняется до тех пор, пока окружающая функция не вернёт управление.


