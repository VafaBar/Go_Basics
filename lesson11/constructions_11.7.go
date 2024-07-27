package main

import "fmt"

func main() {
	heroes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, character := range heroes {
		switch character {
		case 1:
			fmt.Println("Harry")
		case 2:
			fmt.Println("Hermione")
		case 3:
			fmt.Println("Ron")
		case 4:
			fmt.Println("Draco")
			fallthrough
		case 5:
			fmt.Println("Lucius")
		case 6:
			fmt.Println("Albus")
		case 7:
			fmt.Println("Severus")
		case 8:
			fmt.Println("Bellatrix")
		case 9:
			fmt.Println("Sirius")
		case 10:
			fmt.Println("Voldemort")
		default:
			fmt.Println("Unknown character")
		}
	}
}

// Оператор выбора switch с fallthrough (⭐ сложное): Создайте слайс из
// чисел от 1 до 10. Напишите цикл, который использует оператор switch с
// fallthrough для проверки каждого числа и выводит сообщение для
// каждого случая.
