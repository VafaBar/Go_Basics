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
type Order struct {
	Products []Product `json:"products"`
}

func main() {
	pr1 := Product{Name: "HappyCat", Price: 2, Quantity: 200}
	pr2 := Product{Name: "HappyDog", Price: 3, Quantity: 100}
	order := Order{Products: []Product{pr1, pr2}}
	fmt.Println("Структура Order:", order)
	orderJSON, err := json.Marshal(order)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}
	fmt.Println("Сериализованный JSON:", string(orderJSON))

	var decodedOrder Order
	err = json.Unmarshal(orderJSON, &decodedOrder)
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)
		return
	}
	fmt.Println("Десериализованная структура:", decodedOrder)
}

// Сложная структура (⭐ сложное): Создайте структуру Order , которая
// содержит поле Products , являющееся срезом структур Product .
// Сериализуйте и десериализуйте экземпляр структуры Order .
