// Задача 3. Чётность и диапазон
//
// Дано var num int = 17. Проверить:
//
//	isEven    — num чётное (используй % из главы 2.6)
//	inRange   — num строго больше 10 и строго меньше 20
//	isSpecial — верно хотя бы одно из двух условий выше (isEven, inRange)
//
// Ожидаемый вывод:
//
//	isEven: false
//	inRange: true
//	isSpecial: true
//
// Ограничения: isEven и inRange — отдельные bool-переменные; isSpecial
// вычисляется через || именно из isEven и inRange, а не заново из num.
//
// Запуск: go run ./ch02/2.7/task3
package main

import "fmt"

func main() {
	var num int = 17
	var isEven bool = (num % 2) == 0
	var inRange bool = (10 < num) && (20 > num)
	var isSpecial bool = isEven || inRange

	fmt.Printf("isEven: %t\n", isEven)
	fmt.Printf("inRange: %t\n", inRange)
	fmt.Printf("isSpecial: %t\n", isSpecial)

}
