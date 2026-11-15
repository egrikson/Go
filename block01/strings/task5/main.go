// Задача 5. Обрезать, не разрезая букву
//
// Написать функцию truncate(s string, n int) string. Она возвращает самое
// длинное начало строки, которое занимает не больше n байтов и не обрывает
// руну посередине.
//
// Вход и ожидаемый вывод:
//
//	"привет" 5 -> "пр"
//	"привет" 6 -> "при"
//	"привет" 100 -> "привет"
//	"a⌘b" 3 -> "a"
//	"a⌘b" 4 -> "a⌘"
//	"я" 1 -> ""
//	"" 3 -> ""
//
// Ограничения: без []rune. Результат получить срезом строки s[:k], а
// границу найти через индексы for range.
//
// Запуск: go run ./block01/strings/task5
package main

import "fmt"

func main() {
	fmt.Printf("%q 5 -> %q\n", "привет", truncate("привет", 5))
	fmt.Printf("%q 6 -> %q\n", "привет", truncate("привет", 6))
	fmt.Printf("%q 100 -> %q\n", "привет", truncate("привет", 100))
	fmt.Printf("%q 3 -> %q\n", "a⌘b", truncate("a⌘b", 3))
	fmt.Printf("%q 4 -> %q\n", "a⌘b", truncate("a⌘b", 4))
	fmt.Printf("%q 1 -> %q\n", "я", truncate("я", 1))
	fmt.Printf("%q 3 -> %q\n", "", truncate("", 3))
}

func truncate(s string, n int) string {
	end := 0
	if len(s) <= n {
		return s
	}

	for i := range s {
		if i <= n {
			end = i
		}
	}
	return s[:end]
}
