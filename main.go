// Числовые константы
// Числовые константы — это значения высокой точности.

// Нетипизированная константа принимает тип, соответствующий контексту.

// Попробуйте также вывести needInt(Big).

// (В int может храниться максимум 64-битное целое число, а иногда и меньше.
package main

import "fmt"

const (
	// Создаём огромное число, сдвигая единичный бит влево на 100 позиций.
	// Другими словами, это двоичное число, состоящее из 1 и 100 нулей.

	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	// Сдвигаем его ещё раз вправо на 99 позиций, в результате получаем 1 << 1 или 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
// 21
// 0.2
// 1.2676506002282295e+29
