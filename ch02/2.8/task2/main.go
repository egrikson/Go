// Задача 2. Права доступа через битовые флаги
//
// Даны константы-флаги: Read = 1, Write = 2, Execute = 4 (int, каждая —
// отдельный бит). Дано var permissions int = 5. Проверить через & и !=,
// установлен ли у permissions каждый из трёх флагов:
//
//	canRead    — установлен флаг Read
//	canWrite   — установлен флаг Write
//	canExecute — установлен флаг Execute
//
// Ожидаемый вывод:
//
//	canRead: true
//	canWrite: false
//	canExecute: true
//
// Ограничения: Read/Write/Execute — константы; проверка каждого флага —
// через permissions & <флаг> != 0, без сложения или сравнения самого
// числа permissions напрямую с 1/2/4/5.
//
// Запуск: go run ./ch02/2.8/task2
package main

import "fmt"

func main() {
	const (
		Read    int = 1
		Write   int = 2
		Execute int = 4
	)

	var permissions int = 5

	var canRead bool = permissions&Read != 0
	var canWrite bool = permissions&Write != 0
	var canExecute bool = permissions&Execute != 0

	fmt.Printf("canRead: %t\n", canRead)
	fmt.Printf("canWrite: %t\n", canWrite)
	fmt.Printf("canExecute: %t\n", canExecute)
}
