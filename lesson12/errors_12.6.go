package main

import (
	"fmt"
)

type MyError struct {
	Msg            string
	AdditionalInfo string
}

func (e MyError) Error() string {
	return e.Msg + ":" + " " + e.AdditionalInfo
}

func MycustomError(y int) (int, error) {
	if y < 18 {
		return y, MyError{
			Msg:            "minor user",
			AdditionalInfo: "Sorry, but the entrance to the site is strictly from the age of 18",
		}
	}
	return y, nil
}

func main() {
	fmt.Println(MycustomError(16))
}

// Создание кастомного типа ошибки (⭐ сложное): Создайте свой
// собственный тип ошибки, реализовав интерфейс error. Создайте функцию,
// которая возвращает вашу кастомную ошибку, и обработайте её.
