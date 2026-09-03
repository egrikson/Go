// Задача 3. Тарифные лимиты
//
// Объявить блок const:
//
//	BaseLimit = 50
//	BronzeLimit          (без инициализатора, наследует 50)
//	SilverLimit          (без инициализатора, наследует 50)
//	GoldLimit = 200
//	PlatinumLimit        (без инициализатора, наследует 200)
//
// Присвоить значение SilverLimit переменной limit через :=, вывести все
// пять констант, а затем — значение и тип переменной limit:
//
//	BaseLimit: 50
//	BronzeLimit: 50
//	SilverLimit: 50
//	GoldLimit: 200
//	PlatinumLimit: 200
//	limit value: 50
//	limit type: int
//
// Ограничения: BronzeLimit, SilverLimit, PlatinumLimit — без явного
// значения (только автоинициализация); тип переменной limit определяется
// неявно через :=; тип выводится через %T.
//
// Запуск: go run ./ch02/2.4/task3
package main

import "fmt"

func main() {
	const (
		BaseLimit     int = 50
		BronzeLimit       = BaseLimit
		SilverLimit       = BaseLimit
		GoldLimit         = 200
		PlatinumLimit     = GoldLimit
	)

	limit := SilverLimit

	fmt.Println("BaseLimit", BaseLimit)
	fmt.Println("BronzeLimit", BronzeLimit)
	fmt.Println("SilverLimit", SilverLimit)
	fmt.Println("GoldLimit", GoldLimit)
	fmt.Println("PlatinumLimit", PlatinumLimit)
	fmt.Println("limit value", limit)
	fmt.Printf("limit type: %T", limit)

}
