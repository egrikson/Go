// Задача 4. Фильтр — теперь с append
//
// Та же задача, что slices-make/task2, но append уже пройден. Написать
// filterEven(s []int) []int, которая собирает чётные элементы через append,
// начиная с nil-среза. Для входа без чётных функция должна вернуть именно
// nil — проверь это сравнением.
//
// Вход: []int{3, 8, 15, 4, 22, 7, 10} и []int{1, 3, 5}
//
// Ожидаемый вывод:
//
//	Чётные: [8 4 22 10]
//	Из нечётных: [] nil=true
//
// Ограничения: без make; результат начинается с var ... []int; nil=true
// печатается через %t от выражения == nil.
//
// Запуск: go run ./block01/slices-append/task4
package main

import "fmt"

func main() {
	evens := filterEven([]int{3, 8, 15, 4, 22, 7, 10})
	fmt.Println("Чётные:", evens)

	none := filterEven([]int{1, 3, 5})
	fmt.Printf("Из нечётных: %v nil=%t\n", none, none == nil)

}

func filterEven(s []int) []int {
	var result []int
	for _, v := range s {
		if v%2 == 0 {
			result = append(result, v)
		}

	}
	return result
}
