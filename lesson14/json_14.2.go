package main

import (
	"encoding/json"
	"fmt"
)

type Product2 struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

func main() {
	product := Product2{Name: "HappyCat", Price: 2, Quantity: 200}
	fmt.Println("Структура Product2:", product)
	productJSON, err := json.Marshal(product)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}
	fmt.Println("Сериализованный JSON:", string(productJSON))

	var decodedProduct Product2
	err = json.Unmarshal(productJSON, &decodedProduct)
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)
		return
	}
	fmt.Println("Десериализованная структура:", decodedProduct)
}

// Десериализация JSON: Создайте JSON-строку, представляющую данные
// о продукте. Десериализуйте эту строку в структуру Product и выведите
// поля структуры на экран.
