// Задача 1. Максимум
//
// Написать функцию maxOf, принимающую переменное количество int и
// возвращающую наибольшее из них. Если не передано ни одного аргумента —
// вернуть 0.
//
// Вход:
//
//	maxOf(3, 9, 4)
//	maxOf(-5, -2, -9)
//	maxOf(7)
//	maxOf()
//
// Ожидаемый вывод:
//
//	9
//	-2
//	7
//	0
//
// Ограничения: инициализировать «текущий максимум» нулём нельзя — на
// втором вызове получишь 0 вместо -2. Продумай, откуда брать стартовое
// значение и как при этом не выйти за границы при пустом вызове.
// Функция ничего не печатает.
//
// Запуск: go run ./block01/functions-variadic/task1
package main

import "fmt"

func main() {
	fmt.Println(maxOf(3, 9, 4))
	fmt.Println(maxOf(-5, -2, -9))
	fmt.Println(maxOf(7))
	fmt.Println(maxOf())
}

func maxOf(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	var maxValue int = nums[0]
	for _, value := range nums {
		if maxValue < value {
			maxValue = value
		}
	}
	return maxValue
}
