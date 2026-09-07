// Задача 3. Секунды в часы:минуты:секунды
//
// Дано const totalSeconds int = 7325. Разложить это количество секунд
// на часы, минуты и оставшиеся секунды, используя только / и %.
//
// Вывести результат в формате "<часы>ч <минуты>м <секунды>с", затем
// уменьшить секунды на единицу постфиксным декрементом (x--) и вывести
// результат ещё раз в том же формате.
//
// Ожидаемый вывод:
//
//	7325 секунд = 2ч 2м 5с
//	Секундой раньше: 2ч 2м 4с
//
// Ограничения: totalSeconds — const; hours, minutes и секунды —
// обычные переменные, вычисленные через / и %; условные конструкции
// не использовать (тема ещё не пройдена).
//
// Запуск: go run ./ch02/2.6/task3
package main

import "fmt"

func main() {
	const totalSeconds int = 7325
	var hours int = totalSeconds / 3600
	var minutes int = totalSeconds / 60 % 60
	var seconds int = totalSeconds % 60

	fmt.Printf("%d секунд = %dч %dм %dс\n", totalSeconds, hours, minutes, seconds)

	seconds--

	fmt.Printf("Секундой раньше: %dч %dм %dс\n", hours, minutes, seconds)
}
