// Задача 2. Блок и переприсваивание
//
// Объявить три переменные одним блоком var ( ... ): имя, город, возраст.
// Вывести их, затем изменить все три и вывести снова.
//
// Вывести:
//
//	Egor Москва 19
//	Igor Питер 20
//
// Ограничения: обязательно блочное объявление, ровно два вызова
// fmt.Println. Возраст — целое число, тип указать явно.
//
// Запуск: go run ./ch02/2.2/task2
package main

import "fmt"

func main() {
	var (
		name string = "Egor"
		city string = "Москва"
		age  int    = 19
	)

	fmt.Println(name, city, age)

	name = "Igor"
	city = "Питер"
	age = 20

	fmt.Println(name, city, age)
}
