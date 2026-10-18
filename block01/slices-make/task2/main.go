// Задача 2. Фильтр без append
//
// Написать функцию filterEven(s []int) []int, которая возвращает НОВЫЙ
// срез только из чётных элементов s. append ещё не пройден, поэтому
// срез-результат создаётся через make сразу нужной длины.
//
// В main:
//  1. получить чётные из data и напечатать их с len и cap;
//  2. записать -1 в первый элемент результата и напечатать data —
//     исходный срез не должен измениться;
//  3. вызвать filterEven для []int{1, 3, 5} и для nil.
//
// Вход: data := []int{3, 8, 15, 4, 22, 7, 10}
//
// Ожидаемый вывод:
//
//	Чётные: [8 4 22 10] len=4 cap=4
//	Исходный: [3 8 15 4 22 7 10]
//	Из нечётных: [] len=0
//	Из nil: [] len=0
//
// Ограничения: без append; cap результата равен его len (никакого
// запаса «на всякий случай»); исходный срез функция не меняет.
//
// Запуск: go run ./block01/slices-make/task2
//
// НЕ ДОДЕЛАНО: filterEven написана верно, вывод совпадает, но две проверки
// из условия в main не выполняются — совпадение случайное.
//
//  1. Нет пункта 2: перед печатью «Исходный» записать -1 в первый элемент
//     результата. Без этого не проверено, что результат — независимая копия.
//
//  2. Третий вызов идёт с []int{} — это пустой срез, но не nil. Условие
//     просит вызвать filterEven(nil); переменную empty после этого убрать.
package main

import "fmt"

func main() {
	data := []int{3, 8, 15, 4, 22, 7, 10}
	odds := []int{1, 3, 5}
	empty := []int{}

	evens := filterEven(data)
	fromOdds := filterEven(odds)
	fromEmpty := filterEven(empty)

	fmt.Printf("Чётные: %v len=%d cap=%d\n", evens, len(evens), cap(evens))
	fmt.Printf("Исходный: %v\n", data)
	fmt.Printf("Из нечётных: %v len=%d\n", fromOdds, len(fromOdds))
	fmt.Printf("Из nil: %v len=%d\n", fromEmpty, len(fromEmpty))
}

func filterEven(s []int) []int {
	count := 0
	for _, v := range s {
		if v%2 == 0 {
			count++
		}
	}

	evens := make([]int, count)
	next := 0
	for _, v := range s {
		if v%2 == 0 {
			evens[next] = v
			next++
		}
	}
	return evens
}
