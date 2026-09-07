// Задача 1. Все операторы сравнения
//
// Даны var a int = 8, b int = 3. Вывести результат каждого из шести
// операторов сравнения между a и b.
//
// Ожидаемый вывод:
//
//	a == b: false
//	a != b: true
//	a < b: false
//	a > b: true
//	a <= b: false
//	a >= b: true
//
// Ограничения: только int и bool, вывод через fmt.Println.
//
// Запуск: go run ./ch02/2.7/task1
package main

import "fmt"

func main() {
	var a int = 8
	var b int = 3

	fmt.Printf("a == b: %t\n", a == b)
	fmt.Printf("a != b: %t\n", a != b)
	fmt.Printf("a < b: %t\n", a < b)
	fmt.Printf("a > b: %t\n", a > b)
	fmt.Printf("a <= b: %t\n", a <= b)
	fmt.Printf("a >= b: %t\n", a >= b)
}
