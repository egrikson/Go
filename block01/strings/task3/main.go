// Задача 3. Разворот по рунам и по байтам
//
// Написать две функции:
//
//	reverse(s string) string       — разворачивает строку по рунам
//	reverseBytes(s string) string  — разворачивает строку по байтам
//
// Напечатать через %q результаты вызовов в таком порядке:
// reverse для "привет", "go⌘", ""; reverseBytes для "abc", "да".
//
// Ожидаемый вывод:
//
//	"тевирп"
//	"⌘og"
//	""
//	"cba"
//	"\xb0д\xd0"
//
// Ограничения: reverse работает через []rune, reverseBytes через []byte.
// В блоке «Почему» объясни последнюю строку вывода: откуда там \xb0 и \xd0
// и почему буква «д» уцелела.
//
// Запуск: go run ./block01/strings/task3
package main

import "fmt"

func main() {
	listReverse := []string{"привет", "go⌘", ""}
	listReverseBytes := []string{"abc", "да"}

	for _, v := range listReverse {
		fmt.Printf("%q\n", reverse(v))
	}

	for _, v := range listReverseBytes {
		fmt.Printf("%q\n", reverseBytes(v))
	}
}

func reverse(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := range runes[:n/2] {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func reverseBytes(s string) string {
	b := []byte(s)
	n := len(b)
	for i := range b[:n/2] {
		b[i], b[n-1-i] = b[n-1-i], b[i]
	}
	return string(b)
}

// --- Почему ---
//потому что у нас тут кириллица, в которой 1 руна - это 2 байта, а буква Д не изменилась, так как её второй байт равен первому байту А и они тупо склеились
