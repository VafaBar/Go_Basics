package main

import "fmt"

func main() {

	var x [4]int

	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4

	fmt.Println(x)

	x[0] = 4
	x[1] = 3
	x[2] = 2
	x[3] = 1

	fmt.Println(x)

	//  Создайте массив типа int с элементами 1, 2, 3, 4, измените их порядок в массиве на 4, 3, 2, 1 с помощью изменения элементов по индексу.
}