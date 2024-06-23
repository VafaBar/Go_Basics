package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3}

	x := map[string][]int{

		"a": numbers,
		"b": numbers,
		"c": numbers,
	}

	fmt.Println(x)
}

// Создайте map, где значением будут слайс типа int. Создайте несколько пар ключ-значение в данной map. Выведите результат.
