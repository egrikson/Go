// Задача 2. Нулевые значения vs неявный вывод типа
//
// Группа 1 — объявить без инициализации: string, int, bool, float64.
// Группа 2 — объявить через var x = ... (тип выводится неявно):
//
//	целое 7, дробное 3.5, строка "Go", логическое true.
//
// Каждую переменную вывести через:
//
//	fmt.Printf("type: %T value: %v\n", x, x)
//
// Ожидаемый вывод (группа 1, затем группа 2, в этом порядке):
//
//	type: string value:
//	type: int value: 0
//	type: bool value: false
//	type: float64 value: 0
//	type: int value: 7
//	type: float64 value: 3.5
//	type: string value: Go
//	type: bool value: true
//
// Запуск: go run ./ch02/2.3/task2
package main

import "fmt"

func main() {
	var a string
	var b int
	var c bool
	var m float64

	fmt.Printf("type: %T value: %v\n", a, a)
	fmt.Printf("type: %T value: %v\n", b, b)
	fmt.Printf("type: %T value: %v\n", c, c)
	fmt.Printf("type: %T value: %v\n", m, m)

	var b1 = 7
	var m1 = 3.5
	var a1 = "Go"
	var c1 = true

	fmt.Printf("type: %T value: %v\n", b1, b1)
	fmt.Printf("type: %T value: %v\n", m1, m1)
	fmt.Printf("type: %T value: %v\n", a1, a1)
	fmt.Printf("type: %T value: %v\n", c1, c1)
}
