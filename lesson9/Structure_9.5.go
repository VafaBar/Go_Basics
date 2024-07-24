package main

import "fmt"

func main() {

	type Person struct {
		Name      string
		Age       int
		IsMarried bool
	}

	t1 := Person{
		Name:      "Kira",
		Age:       25,
		IsMarried: false,
	}

	t2 := Person{
		Name:      "Anna",
		Age:       25,
		IsMarried: true,
	}

	fmt.Println(t1 == t2)
}

// Сравнение структур: Создайте два экземпляра одной и той же
// структуры. Сравните их с помощью оператора == и выведите результат
