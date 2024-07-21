package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, -6}
	for _, number := range numbers {
		if number > 0 {
			fmt.Println("number is positive")
		}
	}
}

// Условный оператор if: Создайте слайс из чисел. Напишите цикл for-range,
// который проверяет, является ли каждое число положительным, и выводит
// соответствующее сообщение.
