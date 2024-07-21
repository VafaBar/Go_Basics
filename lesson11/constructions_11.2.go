package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	for _, number := range numbers {
		if number > 10 {
			fmt.Println("number is greater than ten")
		} else if number < 10 {
			fmt.Println("number is less than ten")
		} else if number == 10 {
			fmt.Println("number is ten")
		}
	}
}

// Условный оператор if-else: Создайте слайс из чисел. Напишите цикл,
// который проверяет, больше ли каждое число 10, и выводит разные
// сообщения для случаев, когда это верно и когда не верно.
