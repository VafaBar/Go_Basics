package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cat struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Color string `json:"color"`
}

func main() {
	cat := Cat{Name: "Persik", Age: 2, Color: "Ginger"}
	fmt.Println("Структура Cat:", MyJson(cat))
}

func MyJson(cat Cat) string {
	catJSON, err := json.Marshal(cat)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		log.Fatal(err)

	}
	return string(catJSON)
}

// Обработка ошибок сериализации: Напишите функцию, которая
// принимает структуру и возвращает JSON-строку. Обработайте
// возможные ошибки при сериализации и выведите их на экран.
