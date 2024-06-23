package main

import "fmt"

func main() {

	snacks := []string{"Mars", "Snickers", "Bounty", "Twix", "Nuts"}

	snacks = append(snacks[:2], snacks[3:]...)

	fmt.Println(snacks)

	//  Создайте произвольного типа слайс с 5 элементами , с помощью функции append() удалите элемент под индексом 2.
	// Выведите результат на экран.
}
