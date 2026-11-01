// Задача 1. Собрать срез через append
//
// Начать с пустого среза squares и добавить в него квадраты чисел 1..5 —
// по одному за итерацию. Затем одним вызовом append добавить сразу 36 и 49.
// Затем одним вызовом append приклеить срез extra.
//
// Вход: extra := []int{100, 121}
//
// Ожидаемый вывод:
//
//	Квадраты: [1 4 9 16 25]
//	После добавления 36 и 49: [1 4 9 16 25 36 49] len=7
//	Склейка: [1 4 9 16 25 36 49 100 121]
//
// Ограничения: squares объявить как var squares []int (nil), без make
// и без литерала; 36 и 49 — одним append; extra — одним append, не циклом.
//
// Запуск: go run ./block01/slices-append/task1
package main

import "fmt"

func main() {
	extra := []int{100, 121}

	var squares []int

	for i := 1; i <= 5; i++ {
		squares = append(squares, i*i)
	}

	fmt.Println("Квадраты:", squares)

	squares = append(squares, 36, 49)
	fmt.Printf("После добавления 36 и 49: %v len=%d\n", squares, len(squares))

	squares = append(squares, extra...)
	fmt.Println("Склейка:", squares)
}
