package main

import (
	"fmt"
	"math"
)

func main() {

	var d int
	fmt.Print("Введите число: ")
	fmt.Scan(&d)
	a := math.Sqrt(float64(d))

	fmt.Println("Квадратный корень числа: ", a)
}

//  Использование стандартной библиотеки: Импортируйте пакет "math" и
// используйте его для вычисления квадратного корня числа.
