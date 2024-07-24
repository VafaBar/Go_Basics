package main

import (
	"errors"
	"fmt"
)

var Err0 = errors.New("Сan't divide")

func NewError(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("%w: %d and %d", Err0, x, y)

	}
	return x / y, nil

}
func main() {
	fmt.Println(NewError(5, 0))
}

// Использование fmt.Errorf для создания ошибки: Используйте функцию
// fmt.Errorf для создания новой ошибки с форматированным сообщением.
// Верните эту ошибку из вашей функции и обработайте её.
