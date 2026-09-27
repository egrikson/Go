// Задача 1. Минимум и максимум одной функцией
//
// Написать функцию minMax, которая принимает массив [5]int и возвращает
// ДВА значения: минимальный и максимальный элементы.
//
// Вход: [5]int{12, 45, 7, 23, 38}
//
// Ожидаемый вывод:
//
//	Минимум: 7
//	Максимум: 45
//
// Ограничения: одна функция и один проход по массиву — не два цикла
// подряд и не две отдельные функции. Оба значения начинать с первого
// элемента массива. Функция ничего не печатает.
//
// Запуск: go run ./block01/functions-multi/task1
package main

import "fmt"

func main() {
	list := [5]int{12, 45, 7, 23, 38}
	minOf, maxOf := minMax(list)

	fmt.Println("Минимум:", minOf)
	fmt.Println("Максимум:", maxOf)
}

func minMax(data [5]int) (int, int) {
	maxValue := data[0]
	minValue := data[0]

	for _, value := range data {
		if value > maxValue {
			maxValue = value
		}

		if value < minValue {
			minValue = value
		}
	}

	return minValue, maxValue
}
