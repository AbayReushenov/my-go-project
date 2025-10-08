// Константы
// Константы объявляются как переменные, но с ключевым словом const.

// Константы могут иметь символьные, строковые, логические или числовые значения.

// Константы нельзя объявлять с использованием синтаксиса :=.
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
// Hello 世界
// Happy 3.14 Day
// Go rules? true
