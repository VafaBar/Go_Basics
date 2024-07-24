package main

import "fmt"

func main() {
	type Job struct {
		Title  string
		Salary int
	}

	type Employee struct {
		Name   string
		Career Job
	}

	Info := Employee{
		Name: "Steve",
		Career: Job{
			Title:  "economist",
			Salary: 1,
		},
	}
	fmt.Println(Info)
}

// Создание вложенных структур: Создайте структуру "Employee" с полем
// "Name" и вложенной структурой "Job" с полями "Title" и "Salary".
// Создайте экземпляр этой структуры, заполните все поля и выведите их на
// экран.
