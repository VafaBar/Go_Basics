package main

import "fmt"

func main() {
	var ages [5]int
	ages[0] = 12
	ages[1] = 22
	ages[2] = 32
	ages[3] = 42
	ages[4] = 52

	for i, v := range ages {
		fmt.Printf("Индекс: %d, Значение: %d\n", i, v)
	}

}

// Наполнение массива: Создайте массив типа int с 5 элементами, с
// помощью цикла запишите в массив значения, которые будут равняться
// индексу, по которому значение будет находиться
