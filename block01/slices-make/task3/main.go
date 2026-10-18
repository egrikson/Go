// Задача 3. Таблица и транспонирование
//
// Написать функцию makeGrid(rows, cols int, f func(r, c int) int) [][]int:
// создаёт таблицу rows×cols и заполняет каждую ячейку значением f(r, c).
// Написать printGrid(g [][]int), печатающую таблицу построчно.
//
// В main:
//  1. построить таблицу 3×4 с числами от 1 до 12 по строкам;
//  2. построить транспонированную таблицу 4×3 ТОЙ ЖЕ функцией makeGrid:
//     ячейка (r, c) новой таблицы — это ячейка (c, r) исходной;
//  3. напечатать обе.
//
// Ожидаемый вывод:
//
//	Исходная:
//	[1 2 3 4]
//	[5 6 7 8]
//	[9 10 11 12]
//	Транспонированная:
//	[1 5 9]
//	[2 6 10]
//	[3 7 11]
//	[4 8 12]
//
// Ограничения: в main нет циклов — заполнение только через makeGrid, а
// транспонирование — через функцию, переданную в makeGrid (вспомни
// замыкания). Размеры транспонированной таблицы берутся через len из
// исходной, а не числами 4 и 3. Внутри makeGrid — только make, без append.
//
// Запуск: go run ./block01/slices-make/task3
package main

import "fmt"

func main() {
	grid := makeGrid(3, 4, func(r, c int) int {
		return r*4 + c + 1

	})
	fmt.Println("Исходная:")
	printGrid(grid)

	transposed := makeGrid(len(grid[0]), len(grid), func(r, c int) int {
		return grid[c][r]
	})
	fmt.Println("Транспонированная:")
	printGrid(transposed)

}

func makeGrid(rows, cols int, f func(r, c int) int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrixCols := make([]int, cols)
		matrix[i] = matrixCols
		for j := 0; j < cols; j++ {
			matrix[i][j] = f(i, j)
		}
	}
	return matrix
}

func printGrid(g [][]int) {
	for i := 0; i < len(g); i++ {
		fmt.Println(g[i])
	}
}
