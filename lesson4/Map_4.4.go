package main

import "fmt"

func main() {

	x := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	delete(x, "a")

	fmt.Println(x)

	//  Создайте map, добавьте в нее несколько пар ключ-значение. Затем удалите одну из пар с помощью функции delete и выведите всю map на экран.
}
