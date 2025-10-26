package main

import (
	"fmt"
)

// Go не имеет классов
// Однако можно определить методы для типов
// Метод Quality «действует» на экземпляре current типа City,
// используя его поля People, Parks, Cars внутри вычислений.​
type City struct {
	People int32
	Parks float64
	Cars int64
}

func (current City) Quality() float64 {
	return  current.Parks / float64(current.People) + float64(current.Cars) / float64(current.People)
}

func main() {
	Moscow := City{11500000, 100000, 2000000}

	fmt.Println("Качество жизни в Москве:", Moscow.Quality())
}

// aaaaa@aaaaa-GF63-Thin-9SCXR:~/my-go-project$ go run main.go
// Качество жизни в Москве: 0.1826086956521739
