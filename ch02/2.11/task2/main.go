// Задача 2. Месяц: время года и число дней
//
// Дано var month int = 11. Двумя отдельными конструкциями switch вывести
// время года и количество дней в этом месяце.
//
//	Зима   — 12, 1, 2        Весна — 3, 4, 5
//	Лето   — 6, 7, 8         Осень — 9, 10, 11
//
//	31 день — 1, 3, 5, 7, 8, 10, 12
//	30 дней — 4, 6, 9, 11
//	28 дней — 2
//
// Для значений вне 1–12 оба switch должны отработать через default:
// «Такого месяца нет» и 0 дней соответственно.
//
// Ожидаемый вывод:
//
//	Месяц: 11
//	Время года: Осень
//	Дней в месяце: 30
//
// Ограничения: только switch, никаких if; месяцы группировать
// множественными значениями в одном case (case 12, 1, 2:), а не писать
// по case на каждый месяц; в обоих switch обязателен default.
//
// Проверь себя: подставь month = 2 и month = 15 — вывод должен остаться
// осмысленным.
//
// Запуск: go run ./ch02/2.11/task2
package main

import "fmt"

func main() {
	var month int = 11
	var season string
	var dayInMonth int

	switch month {
	case 12, 1, 2:
		season = "Зима"
	case 3, 4, 5:
		season = "Весна"
	case 6, 7, 8:
		season = "Лето"
	case 9, 10, 11:
		season = "Осень"
	default:
		season = "Такого месяца нет"
	}

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		dayInMonth = 31
	case 4, 6, 9, 11:
		dayInMonth = 30
	case 2:
		dayInMonth = 28
	default:
		dayInMonth = 0
	}

	fmt.Println("Месяц:", month)
	fmt.Println("Время года:", season)
	fmt.Println("Дней в месяце:", dayInMonth)

}
