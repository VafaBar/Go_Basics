package main

import (
	"log"
	"os"
)

func main() {

	_, err := os.Open("enigma.ext")
	if err != nil {
		log.Fatal(err)
	}
}

// Возвращение ошибки из функции: Напишите функцию, которая
// возвращает ошибку. Вызовите эту функцию и обработайте возвращенную
// ошибку.
