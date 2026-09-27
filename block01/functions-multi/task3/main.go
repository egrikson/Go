// Задача 3. Три значения и пустой идентификатор
//
// Написать функцию stats(numbers [5]int) (int, int, int), возвращающую
// сумму, минимум и максимум — за один проход по массиву.
//
// Вызвать её дважды:
//  1. взять все три значения и напечатать одной строкой;
//  2. взять ТОЛЬКО максимум, остальные два отбросить через _.
//
// Вход: [5]int{12, 45, 7, 23, 38}
//
// Ожидаемый вывод:
//
//	Сумма: 125, минимум: 7, максимум: 45
//	Только максимум: 45
//
// Ограничения: один цикл внутри stats, не три; первая строка печатается
// одним вызовом Printf, а не тремя Println; во втором вызове ненужные
// значения обязательно через _ — заводить переменные и не использовать
// их нельзя, это ошибка компиляции.
//
// Запуск: go run ./block01/functions-multi/task3
package main

import "fmt"

func main() {
	data := [5]int{12, 45, 7, 23, 38}
	var sumData, maxData, minData = stats(data)
	fmt.Printf("Сумма: %d, минимум: %d, максимум: %d\n", sumData, minData, maxData)
	fmt.Println("Только максимум:", maxData)
}

func stats(numbers [5]int) (int, int, int) {
	var sum int
	var maxValue int
	minValue := numbers[0]

	for _, value := range numbers {
		sum += value
		if value > maxValue {
			maxValue = value
		}
		if value < minValue {
			minValue = value
		}
	}

	return sum, maxValue, minValue
}
