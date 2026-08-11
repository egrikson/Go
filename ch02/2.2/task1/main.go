// Задача 1. Четыре способа объявления
//
// Объявить четыре строковых переменных, каждую своим способом:
//  1. var с типом, значение присвоить отдельной строкой
//  2. var с типом и инициализацией
//  3. var без типа
//  4. краткое объявление :=
//
// Вывести:
//
//	способ один: var с типом
//	способ два: var с типом и значением
//	способ три: var без типа
//	способ четыре: краткое объявление
//
// Ограничения: только fmt.Println, отступ табуляцией.
//
// Запуск: go run ./ch02/2.2/task1
package main

import "fmt"

func main() {
	var hello string
	var name string = "Egor"
	var age = 19
	hobby := "sloboda"

	fmt.Println(hello)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(hobby)
}
