// Задача 1. Побитовые операции над числами
//
// Даны var a int = 12, b int = 10. Вывести a и b в двоичном виде, затем
// результаты (в десятичном и двоичном виде) для a & b, a | b, a ^ b,
// a &^ b, a << 2, a >> 2.
//
// Ожидаемый вывод:
//
//	a = 12 (1100)
//	b = 10 (1010)
//	a & b = 8 (1000)
//	a | b = 14 (1110)
//	a ^ b = 6 (110)
//	a &^ b = 4 (100)
//	a << 2 = 48 (110000)
//	a >> 2 = 3 (11)
//
// Ограничения: вывод через fmt.Printf со спецификаторами %d и %b.
//
// Запуск: go run ./ch02/2.8/task1
package main

import "fmt"

func main() {
	var a int = 12
	var b int = 10

	fmt.Printf("a = %d (%b)\n", a, a)
	fmt.Printf("b = %d (%b)\n", b, b)
	fmt.Printf("a & b = %d (%b)\n", a&b, a&b)
	fmt.Printf("a | b = %d (%b)\n", a|b, a|b)
	fmt.Printf("a ^ b = %d (%b)\n", a^b, a^b)
	fmt.Printf("a &^ b = %d (%b)\n", a&^b, a&^b)
	fmt.Printf("a << 2 = %d (%b)\n", a<<2, a<<2)
	fmt.Printf("a >> 2 = %d (%b)\n", a>>2, a>>2)
}
