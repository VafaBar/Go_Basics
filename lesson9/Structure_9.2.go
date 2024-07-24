package main

import "fmt"

func main() {
	type Person struct {
		Name      string
		Age       int
		IsMarried bool
	}

	myperson := Person{
		Name:      "Vafa",
		Age:       32,
		IsMarried: true,
	}
	fmt.Println(myperson)
}

// Создание структур с разными типами данных: Создайте структуру
// "Person" с полями "Name" (string), "Age" (int) и "IsMarried" (bool). Создайте
// экземпляр этой структуры, заполните поля и выведите их на экран.
