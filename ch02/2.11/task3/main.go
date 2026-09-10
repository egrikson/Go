// Задача 3. Разбор оценок
//
// Дан массив var grades = [5]int{4, 5, 3, 5, 2}. Вывести сам массив,
// максимальную и минимальную оценку, сумму, средний балл (целочисленно)
// и итог по среднему баллу:
//
//	4 и выше — «Отлично»
//	3        — «Нормально»
//	меньше 3 — «Плохо»
//
// Ожидаемый вывод:
//
//	Оценки: [4 5 3 5 2]
//	Максимум: 5
//	Минимум: 2
//	Сумма: 19
//	Средний балл: 3
//	Итог: Нормально
//
// Ограничения: циклов ещё нет, поэтому элементы перебираются вручную —
// но максимум и минимум ДОЛЖНЫ находиться сравнениями через if, а не
// вписываться числом. Проверка простая: если поменять значения в массиве
// на другие, программа обязана дать верный ответ без правок кода.
// Подсказка по приёму: завести переменную с первым элементом и по
// очереди сравнивать её с остальными.
//
// Запуск: go run ./ch02/2.11/task3
package main

import "fmt"

func main() {
	var grades = [5]int{4, 5, 3, 5, 2}

	var maxGrade int = grades[0]
	var minGrade int = grades[0]
	var sumGrades int = grades[0] + grades[1] + grades[2] + grades[3] + grades[4]
	var averageScore int = sumGrades / len(grades)
	var result string

	if maxGrade < grades[1] {
		maxGrade = grades[1]
	}
	if maxGrade < grades[2] {
		maxGrade = grades[2]
	}
	if maxGrade < grades[3] {
		maxGrade = grades[3]
	}
	if maxGrade < grades[4] {
		maxGrade = grades[4]
	}

	if minGrade > grades[1] {
		minGrade = grades[1]
	}
	if minGrade > grades[2] {
		minGrade = grades[2]
	}
	if minGrade > grades[3] {
		minGrade = grades[3]
	}
	if minGrade > grades[4] {
		minGrade = grades[4]
	}

	if averageScore >= 4 {
		result = "Отлично"
	} else if averageScore == 3 {
		result = "Нормально"
	} else {
		result = "Плохо"
	}

	fmt.Println("Оценки:", grades)
	fmt.Println("Максимум:", maxGrade)
	fmt.Println("Минимум:", minGrade)
	fmt.Println("Сумма:", sumGrades)
	fmt.Println("Средний балл:", averageScore)
	fmt.Println("Итог:", result)
}
