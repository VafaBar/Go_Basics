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
	myperson.Name = "Nika"
	fmt.Println(myperson)
}

// Изменение значений полей структуры: Создайте экземпляр структуры
// из предыдущего задания, измените значение одного из полей и выведите
// все поля на экран.
