package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name    string `json:"name"`
	ID      int    `json:"id"`
	Product string `json:"product"`
}

func main() {
	employee := Employee{Name: "Greg", ID: 20202, Product: "HappyCat"}
	fmt.Println("Структура Employee:", employee)

	employeeJSON, err := json.Marshal(employee)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}
	fmt.Println("Сериализованный JSON:", string(employeeJSON))
}

// Использование тегов: Создайте структуру Employee с полями ID , Name , и
// Product . Используйте теги для изменения имен полей в JSON.
// Сериализуйте экземпляр структуры в JSON и убедитесь, что имена полей
// соответствуют тегам.
