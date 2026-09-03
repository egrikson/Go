// Задача 3. Блок var + обмен значениями
//
// В одном блоке var (...) объявить:
//
//	name string = "Alpha"
//	age int32 = 30
//	height float32 = 175.5
//
// Отдельно через := объявить isStudent := true.
//
// Вывести все четыре в формате задачи 2:
//
//	type: string value: Alpha
//	type: int32 value: 30
//	type: float32 value: 175.5
//	type: bool value: true
//
// Завести name2 := "Beta" и поменять name/name2 местами без временной
// переменной (a, b = b, a). Вывести после обмена:
//
//	name: Beta
//	name2: Alpha
//
// Запуск: go run ./ch02/2.3/task3
package main

import "fmt"

func main() {
	var (
		name   string  = "Alpha"
		age    int32   = 30
		height float32 = 175.5
	)

	isStudent := true

	fmt.Printf("type: %T value: %v\n", name, name)
	fmt.Printf("type: %T value: %v\n", age, age)
	fmt.Printf("type: %T value: %v\n", height, height)
	fmt.Printf("type: %T value: %v\n", isStudent, isStudent)

	name2 := "Beta"

	name, name2 = name2, name

	fmt.Println("name:", name)
	fmt.Println("name2:", name2)
}
