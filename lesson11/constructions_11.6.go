package main

import "fmt"

func main() {

	unknowns := []bool{true, false}
	for _, unknown := range unknowns {
		switch unknown {
		case true:

			fmt.Println("YES")

		case false:

			fmt.Println("NO")

		}
	}
}

// Оператор выбора switch без выражения: Создайте слайс из булевых
// значений. Напишите цикл, который использует оператор switch без
// выражения для вывода разных сообщений в зависимости от значения
// булевой переменной.
