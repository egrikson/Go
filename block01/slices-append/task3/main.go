// Задача 3. Добавить в срез из функции
//
// Написать три функции, которые добавляют v в конец среза:
//
//	addBad(s []int, v int)        — внутри s = append(s, v), ничего не возвращает
//	add(s []int, v int) []int     — возвращает результат append
//	addPtr(s *[]int, v int)       — меняет срез через указатель, ничего не возвращает
//
// Вызвать их по очереди для одного и того же среза nums и печатать его
// после каждого вызова.
//
// Вход: nums := []int{1, 2, 3}; addBad(…, 4), add(…, 4), addPtr(…, 5)
//
// Ожидаемый вывод:
//
//	addBad: [1 2 3]
//	add: [1 2 3 4]
//	addPtr: [1 2 3 4 5]
//
// Ограничения: функции именно с такими сигнатурами. В комментарии над
// addBad одной фразой объясни, почему её результат пропадает.
//
// Запуск: go run ./block01/slices-append/task3
//
// НЕ ДОДЕЛАНО: функции и вывод верные, не выполнено одно ограничение.
//
//  1. Над addBad нет комментария, который одной фразой объясняет, почему
//     её результат пропадает. Условие требует его явно. (Линтер ineffassign
//     на строке s = append(s, v) — ожидаем: функция «плохая» по замыслу.)
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}

	addBad(nums, 4)
	fmt.Println("addBad:", nums)

	nums = add(nums, 4)
	fmt.Println("add:", nums)

	addPtr(&nums, 5)
	fmt.Println("addPtr:", nums)

}

func addBad(s []int, v int) {
	s = append(s, v)
}

func add(s []int, v int) []int {
	return append(s, v)
}

func addPtr(s *[]int, v int) {
	*s = append(*s, v)
}
