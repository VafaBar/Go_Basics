package main

import "fmt"

func main() {

	for t := 2; t <= 20; t += 2 {
		fmt.Println(t)
	}

	for d := 0; d <= 20; d++ {
		if d%2 == 0 {
			fmt.Println(d)

		}

	}

}

// Напишите цикл for, который выводит на экран
// все четные числа от 1 до 20.
