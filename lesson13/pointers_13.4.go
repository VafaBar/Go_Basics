package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	myperson := Person{
		Name: "Vafa",
		Age:  32,
	}
	changePerson(&myperson)
}
func changePerson(myperson *Person) {
	myperson.Age = 18
	fmt.Println(myperson)
}

// Указатели на структуры: Создайте структуру и указатель на эту
// структуру. Измените значение поля структуры, используя указатель в
// методе.
