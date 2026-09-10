// Задача 1. Идти гулять или нет
//
// Даны var temp int = -5 и var isWeekend bool = true. Вывести:
//
//  1. температуру;
//  2. категорию погоды через if / else if / else:
//     ниже 0        — «Мороз»
//     от 0 до 14    — «Холодно»
//     от 15 до 24   — «Тепло»
//     25 и выше     — «Жара»
//  3. решение: «Гулять», если сегодня выходной И теплее -10, иначе «Дома».
//
// Ожидаемый вывод:
//
//	Температура: -5
//	Категория: Мороз
//	Решение: Гулять
//
// Ограничения: категория — одна цепочка if / else if / else, без
// повторных проверок уже отсечённых диапазонов (если дошли до второй
// ветки, значит temp уже >= 0 — не надо писать это ещё раз); решение —
// одно условие через &&. Код должен давать верный ответ при любых
// значениях temp и isWeekend, а не только при этих.
//
// Запуск: go run ./ch02/2.11/task1
package main

import "fmt"

func main() {
	var temp int = -5
	var isWeekend bool = true
	var category string
	var solution string

	if temp < 0 {
		category = "Мороз"
	} else if temp <= 14 {
		category = "Холодно"
	} else if temp <= 24 {
		category = "Тепло"
	} else {
		category = "Жара"
	}

	if isWeekend && temp > -10 {
		solution = "Гулять"
	} else {
		solution = "Дома"
	}

	fmt.Println("Температура:", temp)
	fmt.Println("Категория:", category)
	fmt.Println("Решение:", solution)
}
