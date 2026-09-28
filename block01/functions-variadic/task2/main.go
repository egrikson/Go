// Задача 2. Три числа за один проход
//
// Написать функцию stats, принимающую переменное количество int и
// возвращающую ТРИ значения: сумму всех чисел, количество чётных и
// количество отрицательных — именно в таком порядке.
//
// Вход:
//
//	stats(4, -3, 10, -8, 7)
//	stats()
//
// Ожидаемый вывод:
//
//	10 3 2
//	0 0 0
//
// Ограничения: ровно ОДИН цикл по nums — не три прохода подряд.
// Ноль считается чётным и НЕ считается отрицательным.
// Функция ничего не печатает.
//
// Запуск: go run ./block01/functions-variadic/task2
package main

import "fmt"

func main() {
	sum, even, negative := stats(4, -3, 10, -8, 7)
	fmt.Println(sum, even, negative)

	sum, even, negative = stats()
	fmt.Println(sum, even, negative)
}

func stats(nums ...int) (int, int, int) {
	var sumNums, countEven, countNegative int

	for _, value := range nums {
		sumNums += value
		if value%2 == 0 {
			countEven++
		}
		if value < 0 {
			countNegative++
		}
	}
	return sumNums, countEven, countNegative
}
