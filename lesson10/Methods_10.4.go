package main

import "fmt"

type Address struct {
	Street  string
	City    string
	Country string
}

type Person struct {
	Name     string
	Location Address
}

func (p Person) FullAddress() {
	fmt.Printf("%s lives on %s street in %s, %s", p.Name, p.Location.Street, p.Location.City, p.Location.Country)
}

func main() {

	info := Person{
		Name: "Sherlock",
		Location: Address{
			Street:  "Baker",
			City:    "London",
			Country: "Great Britain",
		},
	}
	info.FullAddress()
}

// Создайте структуру "Person" с полем "Name" и вложенной структурой
// "Address" с полями "Street", "City" и "Country". Добавьте метод
// "FullAddress", который возвращает полный адрес как строку.
