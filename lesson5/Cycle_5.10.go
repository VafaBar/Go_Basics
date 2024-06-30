package main

import "fmt"

func main() {

	for t := 1; t < 10; t++ {
		for d := 1; d < 10; d++ {
			fmt.Print(t*d, "\t")
		}
		fmt.Println()
	}

}

//  Цикл в цикле: Используя два цикла выведите таблицу умножения.
