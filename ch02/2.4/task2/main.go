// Задача 2. Дни недели
//
// В одном блоке const объявить дни недели через iota, начиная с Monday (0)
// и заканчивая Sunday (6).
//
// Вывести только Tuesday, Thursday и Saturday:
//
//	Tuesday: 1
//	Thursday: 3
//	Saturday: 5
//
// Ограничения: числовые значения дней получаются только через iota, вручную
// число присваивается только первой константе блока (Monday = iota).
//
// Запуск: go run ./ch02/2.4/task2
package main

import "fmt"

func main() {
	const (
		Monday    int = 0
		Tuesday       = iota
		Wednesday     = iota
		Thursday      = iota
		Friday        = iota
		Saturday      = iota
		Sunday        = iota
	)

	fmt.Println("Tuesday", Tuesday)
	fmt.Println("Thursday", Thursday)
	fmt.Println("Saturday", Saturday)
}
