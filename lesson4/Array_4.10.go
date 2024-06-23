package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := [3]int{}

	a := 34
	b := "88"
	r := 5.68

	x[0] = a

	x[1] = int(r)

	d, _ := strconv.Atoi(b)

	x[2] = d

	fmt.Println(x)

	// Создайте массив типа int с 3 элементами, также создайте 3 переменные разных типов (int, string, float64).
	// Сохраните элементы массива в переменные с помощью конвертации. Выведите результат.
}
