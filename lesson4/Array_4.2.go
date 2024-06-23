package main

import "fmt"

func main() {
	var colors [3]string

	colors[0] = "green"
	colors[1] = "pink"
	colors[2] = "red"

	colors[1] = "blue"

	fmt.Println(colors)
	//Создайте массив из трех строк. Измените значение второго элемента и выведите все элементы на экран.
}
