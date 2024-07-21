package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140}
	for _, number := range numbers {
		if number == 10 {
			fmt.Println("number is true")
		} else if number == 20 {
			fmt.Println("number is true")
		} else if number == 30 {
			fmt.Println("number is true")
		} else {
			fmt.Println("number is false")
		}
	}
}

// Условный оператор if-else: Создайте слайс из чисел. Напишите цикл,
// который сравнивает каждое число со значениями 10, 20 и 30, и выводит
// разные сообщения в зависимости от результата сравнения.
