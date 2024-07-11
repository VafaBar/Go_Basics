package main

import "fmt"

type Book struct {
	Title string
	Page  int
}

func main() {

	r1 := Book{
		Title: "Murder on the Orient Express",
		Page:  302,
	}

	r2 := Book{
		Title: "Death on the Nile",
		Page:  304,
	}

	fmt.Println(Pages(r1, r2).Title)

}
func Pages(t1 Book, t2 Book) Book {
	if t1.Page < t2.Page {
		return t2
	} else {
		return t1
	}

}

// Структуры и функции (⭐ сложное): Создайте функцию, которая
// принимает два аргумента типа "Book" и возвращает книгу с большим
// количеством страниц. Протестируйте эту функцию с несколькими
// экземплярами структуры "Book".
