package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

func main() {
	product := Product{Name: "HappyCat", Price: 2, Quantity: 200}
	fmt.Println("Структура Product:", product)

	productJSON, err := json.Marshal(product)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}
	fmt.Println("Сериализованный JSON:", string(productJSON))
}

// Сериализация структуры: Создайте структуру Product с полями Name ,
// Price и Quantity . Создайте экземпляр этой структуры и сериализуйте его в
// JSON.
