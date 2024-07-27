package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	for _, number := range numbers {
		switch number % 2 {
		case 0:

			fmt.Println("number is even")

		case 1:

			fmt.Println("number is ood")

		}
	}
}

// Оператор выбора switch с несколькими значениями для case: Создайте
// слайс из чисел. Напишите цикл, который использует оператор switch для
// проверки каждого числа и выводит сообщение для четных и нечетных
// чисел.
