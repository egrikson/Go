// Задача 1. Байты и руны
//
// Написать функцию runeCount(s string) int, которая считает руны в строке.
// Для каждой строки из входа напечатать её в кавычках, число байтов и
// число рун.
//
// Вход: "hello", "привет", "go⌘", ""
//
// Ожидаемый вывод:
//
//	"hello": байт=5 рун=5
//	"привет": байт=12 рун=6
//	"go⌘": байт=5 рун=3
//	"": байт=0 рун=0
//
// Ограничения: руны считать через for range по строке, без []rune.
//
// Запуск: go run ./block01/strings/task1
package main

import "fmt"

func main() {
	list := []string{"hello", "привет", "go⌘", ""}
	for _, v := range list {
		fmt.Printf("%q: байт=%d рун=%d\n", v, len(v), runeCount(v))
	}
}

func runeCount(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}
