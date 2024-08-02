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
	fmt.Println("JSON Cat:", MyJson(cat))
	catJ := MyJson(cat)
	fmt.Println("Структура Cat:", DecodedJson(catJ))

}

func MyJson(cat Cat) string {

	catJSON, err := json.Marshal(cat)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		log.Fatal(err)
	}
	return string(catJSON)
}
func DecodedJson(cat string) Cat {
	var decodedCat Cat
	err := json.Unmarshal([]byte(cat), &decodedCat)
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)

	}
	return decodedCat

}

// Обработка ошибок десериализации: Напишите функцию, которая
// принимает JSON-строку и возвращает структуру. Обработайте
// возможные ошибки при десериализации и выведите их на экран.
