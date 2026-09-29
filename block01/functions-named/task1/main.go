// Задача 1. divmod с именами
//
// Переписать функцию divmod из block01/functions-multi/task4 так, чтобы
// возвращаемые значения были ИМЕНОВАННЫМИ:
//
//	func divmod(a, b int) (quotient, remainder int)
//
// Внутри присвоить значения этим переменным и выйти голым return —
// без перечисления значений после него.
//
// Вызвать дважды: divmod(3725, 3600), затем divmod от полученного
// остатка и 60.
//
// Ожидаемый вывод:
//
//	3725 / 3600 = 1, остаток 125
//	125 / 60 = 2, остаток 5
//
// Ограничения: ровно один голый return в конце функции, никаких
// return quotient, remainder. Функция не печатает.
//
// Сравни сигнатуру с прежней версией — (int, int) против
// (quotient, remainder int). Во второй видно, что первым идёт частное,
// и лезть в тело функции, чтобы это выяснить, не нужно.
//
// Запуск: go run ./block01/functions-named/task1
package main

import "fmt"

func main() {
	const total = 3725

	hours, rest := divmod(total, 3600)
	fmt.Printf("%d / 3600 = %d, остаток %d\n", total, hours, rest)

	minutes, seconds := divmod(rest, 60)
	fmt.Printf("%d / 60 = %d, остаток %d\n", rest, minutes, seconds)
}

func divmod(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b

	return
}
