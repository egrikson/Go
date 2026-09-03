// Задача 1. Уровни доступа
//
// Объявить в одном блоке const три именованные константы типа int:
//
//	Guest = 1, User = 2, Admin = 3
//
// Вывести через fmt.Println("Имя:", значение) в этом порядке:
//
//	Guest: 1
//	User: 2
//	Admin: 3
//
// Ограничения: все константы — в одном блоке const, тип int указан явно
// для каждой.
//
// Запуск: go run ./ch02/2.4/task1
package main

import "fmt"

func main() {
	const (
		Guest int = 1
		User int = 2
		Admin int = 3
	)

	fmt.Println("Guest:", Guest)
	fmt.Println("User:", User)
	fmt.Println("Admin:", Admin)
}
