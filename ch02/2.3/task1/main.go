// Задача 1. Все целочисленные типы
//
// Объявить переменные всех типов из таблицы главы:
//
//	int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint
//
// Значения по порядку: 100, 200, 30000, 60000, 2000000000, 4000000000,
// 5000000000000, 10000000000000, 42, 42.
//
// Вывести через fmt.Println("тип:", переменная) в этом же порядке:
//
//	int8: 100
//	uint8: 200
//	int16: 30000
//	uint16: 60000
//	int32: 2000000000
//	uint32: 4000000000
//	int64: 5000000000000
//	uint64: 10000000000000
//	int: 42
//	uint: 42
//
// Запуск: go run ./ch02/2.3/task1
package main

import "fmt"

func main() {
	var a int8 = 100
	var b uint8 = 200
	var c int16 = 30000
	var d uint16 = 60000
	var f int32 = 2000000000
	var s uint32 = 4000000000
	var m int64 = 5000000000000
	var j uint64 = 10000000000000
	var i int = 42
	var y uint = 42

	fmt.Println("int8:", a)
	fmt.Println("uint8:", b)
	fmt.Println("int16:", c)
	fmt.Println("uint16:", d)
	fmt.Println("int32:", f)
	fmt.Println("uint32:", s)
	fmt.Println("int64:", m)
	fmt.Println("uint64:", j)
	fmt.Println("int:", i)
	fmt.Println("uint:", y)
}
