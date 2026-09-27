// Задача 2. Функция вызывает функцию
//
// Дан массив var data = [5]int{12, 45, 7, 23, 38}. Написать три функции:
//
//	sum(numbers [5]int) int     — сумма элементов
//	maxOf(numbers [5]int) int   — максимальный элемент
//	average(numbers [5]int) int — среднее (целочисленно)
//
// Ожидаемый вывод:
//
//	Сумма: 125
//	Максимум: 45
//	Среднее: 25
//
// Ограничения: ни одна из трёх функций не печатает — печатает main.
// average НЕ считает сумму заново: она вызывает sum и делит результат
// на длину массива. Максимум начинать с первого элемента, а не с нуля.
//
// Смысл задачи — увидеть, что возвращённое значение можно сразу
// передать дальше: функция становится кирпичом, из которого собираются
// другие функции. С печатающими функциями так сделать было нельзя.
//
// Имя maxOf, а не max: max — встроенная функция Go, своё имя её затенит.
//
// Запуск: go run ./block01/functions-return/task2
package main

import "fmt"

func main() {
	var data = [5]int{12, 45, 7, 23, 38}

	fmt.Println("Сумма:", sum(data))
	fmt.Println("Максимум:", maxOf(data))
	fmt.Println("Среднее:", average(data))
}

func sum(numbers [5]int) int {
	var total int

	for _, value := range numbers {
		total += value
	}
	return total
}

func maxOf(numbers [5]int) int {
	maxValue := numbers[0]

	for _, value := range numbers {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

func average(numbers [5]int) int {
	return sum(numbers) / len(numbers)
}
