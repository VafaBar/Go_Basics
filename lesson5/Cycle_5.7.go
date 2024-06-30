package main

import "fmt"

func main() {

	colors := []string{"green", "red", "pink"}
	for _, color := range colors {
		fmt.Println(color)
	}
}

// Цикл for-range с слайсом: Создайте слайс с несколькими элементами.
// Напишите цикл for-range, который выводит каждый элемент слайса на
// экран.
