package main

import "fmt"

func main() {
	var names [5]string

	names[0] = "Harry"
	names[1] = "Ron"
	names[2] = "Hermione"
	names[3] = "Draco"
	names[4] = "Polumna"

	for _, name := range names {

		fmt.Println(name)
	}

}

// Цикл for-range с массивом: Создайте массив с несколькими элементами.
// Напишите цикл for-range, который выводит каждый элемент массива на
// экран.
