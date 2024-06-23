package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4}

	y := make([]int, 4)
	copy(y, x)

	fmt.Println(x, y)

	// Создайте два слайса, скопируйте элементы из одного слайса в другой
	// с помощью функции copy и выведите оба слайса на экран.
}
