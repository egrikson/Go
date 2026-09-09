// Задача 3. Сдвиги вместо умножения, деления и остатка
//
// Дано var value int = 37. Вычислить:
//
//	doubled    — value умноженное на 2, через <<
//	quadrupled — value умноженное на 4, через <<
//	halved     — value делённое на 2 (целочисленно), через >>
//	lastBit    — младший бит value (то же, что value % 2), через &
//
// Ожидаемый вывод:
//
//	doubled: 74
//	quadrupled: 148
//	halved: 18
//	lastBit: 1
//
// Ограничения: doubled, quadrupled и halved — только через << и >>, без
// *, / и %; lastBit — только через &, без %.
//
// Запуск: go run ./ch02/2.8/task3
package main

import "fmt"

func main() {
	var value int = 37

	fmt.Printf("doubled: %d\n", value<<1)
	fmt.Printf("quadrupled: %d\n", value<<2)
	fmt.Printf("halved: %d\n", value>>1)
	fmt.Printf("lastBit: %d", value&1)
}
