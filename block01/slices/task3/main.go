// Задача 3. Преобразование на месте
//
// Написать функцию forEach(s []int, f func(int) int), которая заменяет
// каждый элемент среза на f(элемент).
//
//  1. Применить к arr[0:5] с функцией возведения в квадрат.
//  2. Применить к arr[3:8] с функцией, которая прибавляет к элементу
//     номер своего вызова: при первом вызове +1, при втором +2 и т.д.
//  3. В конце напечатать первый срез.
//
// Вход: arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
//
// Ожидаемый вывод:
//
//	После квадратов: [1 4 9 16 25 6 7 8]
//	После прибавки: [1 4 9 17 27 9 11 13]
//	Первый срез: [1 4 9 17 27]
//
// Ограничения: оба среза создаются в переменных ДО первого вызова
// forEach, и последняя строка печатает тот самый первый срез. Счётчик
// вызовов живёт в замыкании — без глобальной переменной и без лишнего
// параметра. forEach ничего не знает о том, какое преобразование ей
// передали.
//
// Запуск: go run ./block01/slices/task3
package main

import (
	"fmt"
)

func main() {
	arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	sliceSquare := arr[0:5]
	sliceMakeAdder := arr[3:8]

	forEach(sliceSquare, square)
	fmt.Printf("После квадратов: %v\n", arr)

	forEach(sliceMakeAdder, makeAdder())
	fmt.Printf("После прибавки: %v\n", arr)

	fmt.Printf("Первый срез: %v\n", sliceSquare)
}

func makeAdder() func(int) int {
	count := 0
	return func(x int) int {
		count++
		return x + count
	}
}

func square(n int) int {
	return n * n
}

func forEach(s []int, f func(int) int) {
	for i, value := range s {
		s[i] = f(value)
	}
}
