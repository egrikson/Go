// Задача 4*. Строго по возрастанию — без цепочек сравнений
//
// Даны var x int = 5, y int = 10, z int = 15. Проверить:
//
//	isIncreasing — x < y и y < z одновременно
//	isDecreasing — x > y и y > z одновременно
//
// Ожидаемый вывод:
//
//	isIncreasing: true
//	isDecreasing: false
//
// Ограничения: писать конструкцию вида x < y < z нельзя — в Go это не
// компилируется (x < y даёт bool, а bool < z — ошибка типов, сравнивать
// с числом нельзя). Каждая проверка — одно составное bool-выражение
// через &&, без промежуточных переменных.
//
// Запуск: go run ./ch02/2.7/task4
package main

import "fmt"

func main() {
	var x int = 5
	var y int = 10
	var z int = 15

	var isIncreasing bool = (x < y) && (y < z)
	var isDecreasing bool = (x > y) && (y > z)

	fmt.Printf("isIncreasing: %t\n", isIncreasing)
	fmt.Printf("isDecreasing: %t\n", isDecreasing)
}
