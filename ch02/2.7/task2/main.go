// Задача 2. Логические операторы
//
// Даны var age int = 20, hasTicket bool = true, price int = 150,
// budget int = 200. Проверить:
//
//	canEnter  — возраст не меньше 18 И есть билет (age, hasTicket)
//	canAfford — бюджета достаточно, чтобы позволить себе цену (budget, price)
//	canGo     — верны оба условия одновременно (canEnter, canAfford)
//
// Ожидаемый вывод:
//
//	canEnter: true
//	canAfford: true
//	canGo: true
//
// Ограничения: canEnter, canAfford, canGo — отдельные bool-переменные;
// canGo вычисляется через && из canEnter и canAfford, а не пересчитывается
// заново из age/hasTicket/price/budget.
//
// Запуск: go run ./ch02/2.7/task2
package main

import "fmt"

func main() {
	var age int = 20
	var hasTicket bool = true
	var price int = 150
	var budget int = 200

	var canEnter bool = (age >= 18) && hasTicket
	var canAfford bool = price <= budget
	var canGo bool = canEnter && canAfford

	fmt.Printf("canEnter: %t\n", canEnter)
	fmt.Printf("canAfford: %t\n", canAfford)
	fmt.Printf("canGo: %t\n", canGo)
}
