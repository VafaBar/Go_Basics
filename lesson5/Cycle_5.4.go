package main

import "fmt"

func main() {
	g := 0
	for d := 1; d <= 100; d++ {
		g = g + d // g += d
		if g >= 50 {
			break
		}
		fmt.Println(d)

	}

}

// Цикл for с break: Напишите цикл for, который выводит числа от 1 до 100, но
// прерывает цикл, как только сумма выводимых чисел превысит 50.
