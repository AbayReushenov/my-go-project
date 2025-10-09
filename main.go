package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	// 42
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i
	// 21

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
	// 73
}


// Указатели
// В Go есть указатели. Указатель хранит адрес памяти значения.

// Тип *T — это указатель на значение T. Его нулевое значение — nil.

// var p *int
// Оператор & генерирует указатель на свой операнд.

// i := 42
// p = &i
// Оператор * обозначает базовое значение указателя.

// fmt.Println(*p) // прочитать i по указателю p
// *p = 21 // установить i по указателю p
// Это называется «разыменованием» или «косвенным обращением».

// В отличие от C, в Go нет арифметики указателей.
