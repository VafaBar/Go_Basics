package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}
type Company struct {
	Employees []Employee `json:"employees"`
}

func main() {
	emp1 := Employee{Name: "Shelby", ID: 26058}
	emp2 := Employee{Name: "Samanta", ID: 31967}
	company := Company{Employees: []Employee{emp1, emp2}}
	fmt.Println("Структура Company:", company)
	companyJSON, err := json.Marshal(company)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}
	fmt.Println("Сериализованный JSON:", string(companyJSON))

	var decodedCompany Company
	err = json.Unmarshal(companyJSON, &decodedCompany)
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)
		return
	}
	fmt.Println("Десериализованная структура:", decodedCompany)
}

// Вложенные структуры (⭐ сложное): Создайте структуру Company , которая
// содержит поле Employees , являющееся срезом структур Employee.
// Сериализуйте и десериализуйте экземпляр структуры Company.
