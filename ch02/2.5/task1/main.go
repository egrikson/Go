// Задача 1. Значение и тип
//
// Даны переменные:
//
//	var name string = "Egor"
//	var age int = 19
//	var height float64 = 183.5
//	var isStudent bool = true
//
// Вывести для каждой переменной строку вида "имя = значение (тип)",
// используя fmt.Printf со спецификаторами %v и %T.
//
// Ожидаемый вывод:
//
//	name = Egor (string)
//	age = 19 (int)
//	height = 183.5 (float64)
//	isStudent = true (bool)
//
// Ограничения: вывод — только через fmt.Printf, ровно 4 вызова.
//
// Запуск: go run ./ch02/2.5/task1
package main

import "fmt"

func main() {
	var name string = "Egor"
	var age int = 19
	var height float64 = 183.5
	var isStudent bool = true

	fmt.Printf("name = %v (%T)\n", name, name)
	fmt.Printf("age = %v (%T)\n", age, age)
	fmt.Printf("height = %v (%T)\n", height, height)
	fmt.Printf("isStudent = %v (%T)\n", isStudent, isStudent)
}
