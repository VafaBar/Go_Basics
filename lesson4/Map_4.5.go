package main

import "fmt"

func main() {

	x := map[string]int{

		"a": 1,
		"b": 2,
		"c": 3,
	}

	if _, ok := x["b"]; ok {
		fmt.Println(ok)
	} else {
		fmt.Println(ok)
	}

	//  Создайте map, добавьте в нее несколько пар ключ-значение. Затем проверьте существование
	// одного из ключей с помощью двухзначной формы индексирования и выведите результат на экран.
}
