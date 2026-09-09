// Задача 4*. Один оператор, без скобок и лишних переменных
//
// Даны var price int = 250, var tax int = 8, var fee int = 12. Нужно
// увеличить price на (tax процентов от price, отброшенное вниз) плюс
// fee — ровно одной строкой через +=.
//
// Ожидаемый вывод:
//
//	price: 282
//
// Ограничения: ровно одна строка ВЫЧИСЛЕНИЯ (печать не в счёт); никаких
// скобок; никаких дополнительных переменных; только price, tax, fee и
// оператор +=.
//
// Запуск: go run ./ch02/2.9/task4
package main

import "fmt"

func main() {
	var price int = 250
	var tax int = 8
	var fee int = 12

	price += price*tax/100 + fee

	fmt.Printf("price: %d", price)
}
