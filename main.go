package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)


// a len=5 cap=5 [0 0 0 0 0]
// b len=0 cap=5 []
// c len=2 cap=5 [0 0]
// d len=3 cap=3 [0 0 0]
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
// Создание среза с помощью make
// Срезы можно создавать с помощью встроенной функции make; именно так создаются массивы с динамическим размером.

// Функция make выделяет обнулённый массив и возвращает срез, ссылающийся на этот массив:

// a := make([]int, 5) // len(a)=5
// Чтобы указать ёмкость, передайте make третий аргумент:

// b := make([]int, 0, 5) // len(b)=0, cap(b)=5

// b = b[:cap(b)] // len(b)=5, cap(b)=5
// b = b[1:] // len(b)=4, cap(b)=4


