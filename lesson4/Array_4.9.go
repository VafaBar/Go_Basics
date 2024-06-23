package main

import "fmt"

func main() {

	fruits1 := "banana"
	fruits2 := "apple"
	fruits3 := "plum"

	var fruits [3]string

	fruits[0] = fruits1
	fruits[1] = fruits2
	fruits[2] = fruits3

	fmt.Println(fruits)

	//  Создайте 3 переменных с разными значениями типа string, создайте массив типа string с 3 элементами, положите в массив значения переменных под разные индексы.
}
