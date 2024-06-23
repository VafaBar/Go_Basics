package main

import "fmt"

func main() {

	list := [5]int{}
	list[0] = 10
	list[1] = 20
	list[2] = 30
	list[3] = 40
	list[4] = 50

	slice := []int{list[0], list[1], list[2]}
	slice = append(slice, 9)

	newslice := make([]int, 4)
	copy(newslice, slice)
	newslice = append(newslice, 8, 7)

	fmt.Println(list, slice, newslice)

	// Создайте массив типа int с 5 элементами 1, 2 3, 4, 5, после чего создайте слайс на основе
	// первых трёх элементов созданного массива с помощью оператора слайса, добавьте к нему элемент 9 с помощью функции append(),
	// после чего создайте новый слайс на основе первого и добавьте к нему два элемента 8 и 7, выведите массив и оба слайса на экран.
	// Попробуйте объяснить полученный результат.

}
