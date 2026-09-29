// Задача 4*. Сводка по переменному числу чисел
//
// Написать функцию с вариативным параметром и ЧЕТЫРЬМЯ именованными
// возвратами:
//
//	func summarize(nums ...int) (min, max, avg int, ok bool)
//
// Для непустого вызова вернуть минимум, максимум, среднее
// (целочисленное) и true. Для пустого вызова — нули и false.
//
// Вход:
//
//	summarize(4, -3, 10, -8, 7)
//	summarize()
//
// Ожидаемый вывод:
//
//	min=-8 max=10 avg=2
//	пустой вызов: данных нет
//
// Ограничения: пустой случай обработать ГОЛЫМ return — здесь он как
// раз уместен; min и max начинать с nums[0]; сумму считать во
// вспомогательной переменной, а не в avg; оба результата проверять
// краткой конструкцией if.
//
// Почему тут проверка на пустоту НУЖНА, а в вариативной задаче 2 была
// лишней. Там цикл по пустому набору просто не выполнялся, и нули
// оказывались верным ответом. Здесь иначе: сразу после проверки идёт
// nums[0], а обращение к первому элементу пустого набора уронит
// программу. Защищать надо то, что действительно падает.
//
// И обрати внимание, как красиво ложится голый return в пустом случае:
// min, max, avg уже нули, ok уже false — ровно то, что нужно вернуть.
// Писать return 0, 0, 0, false не требуется.
//
// Проверь себя: summarize(5) должно дать min=5 max=5 avg=5 ok=true.
//
// Запуск: go run ./block01/functions-named/task4
package main

import "fmt"

func main() {
	if minNum, maxNum, avgNum, okNum := summarize(4, -3, 10, -8, 7); okNum {
		fmt.Printf("min=%d max=%d avg=%d\n", minNum, maxNum, avgNum)
	} else {
		fmt.Printf("пустой вызов: данных нет\n")
	}

	if minNum, maxNum, avgNum, okNum := summarize(); okNum {
		fmt.Printf("min=%d max=%d avg=%d\n", minNum, maxNum, avgNum)
	} else {
		fmt.Printf("пустой вызов: данных нет\n")
	}
}

func summarize(nums ...int) (min, max, avg int, ok bool) {
	if len(nums) == 0 {
		return
	}

	sumNums := 0
	min = nums[0]
	max = nums[0]
	for _, value := range nums {
		sumNums += value
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	avg = sumNums / len(nums)
	ok = true
	return
}
