// Задача 1. Первая своя программа
//
// Вывести три строки:
//
//	Привет, Go!
//	Версия: 1.26.4
//	Учусь по metanit
//
// Ограничения: только fmt.Println, ровно три вызова.
// Форматирование по gofmt: пустая строка после package main, отступ табуляцией.
//
// Запуск: go run ./ch01/1.2/task1
package main

import "fmt"

func main() {
	fmt.Println("Привет, Go!")
	fmt.Println("Версия 1.26.4")
	fmt.Println("Учусь по metanit")
}
