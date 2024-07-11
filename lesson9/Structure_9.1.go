package main

import "fmt"

func main() {
	type Book struct {
		Title  string
		Author string
		Pages  int
	}
	t := Book{
		Title:  "Murder on the Orient Express",
		Author: "Agatha Christie",
		Pages:  302,
	}
	fmt.Println(t)

}

// Создание и использование структур: Создайте структуру,
// представляющую собой "Book" с полями "Title", "Author" и "Pages".
// Создайте экземпляр этой структуры, заполните поля и выведите их на
// экран.
