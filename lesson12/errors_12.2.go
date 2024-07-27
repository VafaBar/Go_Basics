package main

import (
	"errors"
	"fmt"
)

func main() {
	e := errorchik()
	if e != nil {
		fmt.Println(e)
	}
}
func errorchik() error {

	err := errors.New("oh my code")
	return err
}

// Создание новой ошибки: Используйте функцию errors.New для создания
// новой ошибки. Верните эту ошибку из вашей функции и обработайте её.
