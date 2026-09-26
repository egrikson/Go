// Задача 3. Одна функция вызывает другую
//
// Дан массив продаж по дням: [5]int{3, 7, 5, 2, 6}. Построить текстовую
// диаграмму: для каждого дня строка «День N: » и столько звёздочек,
// сколько продаж.
//
// Разбить на ДВЕ функции:
//
//	printBar(day int, value int)  — печатает одну строку диаграммы
//	printChart(data [5]int)       — обходит массив и вызывает printBar
//
// Ожидаемый вывод:
//
//	День 1: ***
//	День 2: *******
//	День 3: *****
//	День 4: **
//	День 5: ******
//
// Ограничения: main вызывает только printChart — напрямую печатать в
// main нельзя. printBar ничего не знает про массив: она получает номер
// дня и число, и всё. Нумерация дней с единицы, а индексы в массиве
// с нуля — это несовпадение решать в printChart, а не внутри printBar.
//
// Подсказка по звёздочкам: fmt.Print печатает без переноса строки,
// fmt.Println() с пустыми скобками переносит строку.
//
// Запуск: go run ./block01/functions/task3
package main

import (
	"fmt"
	"strings"
)

func main() {
	list_bay := [5]int{3, 7, 5, 2, 6}
	printChart(list_bay)
}

func printBar(day int, value int) {
	fmt.Printf("День %d: ", day)
	fmt.Println(strings.Repeat("*", value))
}

func printChart(data [5]int) {
	for i := 0; i < len(data); i++ {
		printBar(i+1, data[i])
	}
}
