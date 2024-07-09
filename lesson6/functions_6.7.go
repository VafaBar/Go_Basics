package main

import "fmt"

func main() {
	sum(1, 2, 3, 4)
}

func sum(nums ...int) {
	total := 0
	d := len(nums)
	for _, num := range nums {
		total += num
	}
	fmt.Println(total, d)
}

// Функция, возвращающая сумму и количество аргументов: Создайте
// функцию, которая получает переменное количество аргументов типа int,
// складывает их, и возвращает их сумму и количество полученных ею
// аргументов.
