// Задача 3. Символ и статусы
//
// Часть А. Дана переменная var letter rune = 'A'. Одним вызовом fmt.Printf
// вывести её в семи представлениях подряд через пробел: десятичный код,
// сам символ, символ в кавычках, шестнадцатеричное, восьмеричное,
// двоичное представление и код в формате Unicode.
//
// Часть Б. В одном блоке const с использованием iota объявить три
// константы без явного типа: StatusNew, StatusActive, StatusDone
// (значения по порядку 0, 1, 2). Вывести для каждой строку вида
// "имя = значение (тип)" через %d и %T.
//
// Ожидаемый вывод:
//
//	65 A 'A' 41 101 1000001 U+0041
//	StatusNew = 0 (int)
//	StatusActive = 1 (int)
//	StatusDone = 2 (int)
//
// Ограничения: часть А — ровно один вызов fmt.Printf; часть Б — все три
// константы в одном блоке const, значение указано только у первой.
//
// Запуск: go run ./ch02/2.5/task3
package main

import "fmt"

func main() {
	var letter rune = 'A'
	fmt.Printf("%v %c %q %x %o %b %U\n", letter, letter, letter, letter, letter, letter, letter)

	const (
		StatusNew = iota
		StatusActive
		StatusDone
	)
	fmt.Printf("StatusNew = %d (%T)\n", StatusNew, StatusNew)
	fmt.Printf("StatusActive = %d (%T)\n", StatusActive, StatusActive)
	fmt.Printf("StatusDone = %d (%T)\n", StatusDone, StatusDone)
}
