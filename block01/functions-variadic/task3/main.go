// Задача 3. Склейка через разделитель
//
// Написать две функции:
//
//	join(sep string, parts ...string) string
//	path(parts ...string) string
//
// join склеивает части через разделитель: разделитель ставится МЕЖДУ
// элементами, в начале и в конце его быть не должно.
//
// path возвращает то же самое с разделителем "/" и обязана вызывать
// join, передавая ей свой параметр через распаковку (parts...).
//
// Вход:
//
//	join(", ", "a", "b", "c")
//	join("-", "one")
//	join("+")
//	path("usr", "local", "bin")
//
// Ожидаемый вывод (третья строка — пустая):
//
//	a, b, c
//	one
//
//	usr/local/bin
//
// Ограничения: пакет strings не использовать, склеивать только
// оператором +. В path не дублировать логику склейки — тело в одну
// строку. Функции ничего не печатают.
//
// Запуск: go run ./block01/functions-variadic/task3
package main

import "fmt"

func main() {
	fmt.Println(join(", ", "a", "b", "c"))
	fmt.Println(join("-", "one"))
	fmt.Println(join("+"))
	fmt.Println(path("usr", "local", "bin"))
}

func join(sep string, parts ...string) string {
	var stringJoin string
	for i, value := range parts {
		if i > 0 {
			stringJoin += sep
		}
		stringJoin += value
	}
	return stringJoin
}

func path(parts ...string) string {
	return join("/", parts...)
}
