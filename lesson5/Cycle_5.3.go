package main

import "fmt"

func main() {

	for d := 0; d <= 20; d++ {
		if d%2 == 1 {
			continue
		}
		fmt.Println(d)

	}

}

// Цикл for с continue: Модифицируйте предыдущий цикл так, чтобы он
// использовал команду continue для пропуска нечетных чисел.
