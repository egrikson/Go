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

// Бесполезна: append возвращает новый описатель среза, а присваивание кладёт
// его в локальную копию параметра — наружу не уходит ни новая длина, ни
// указатель на новый массив, если append его перевыделил.
func addBad(s []int, v int) {
	s = append(s, v)
}

func add(s []int, v int) []int {
	return append(s, v)
}

func addPtr(s *[]int, v int) {
	*s = append(*s, v)
}
