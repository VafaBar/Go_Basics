package main

import (
	"errors"
	"fmt"
)

var Err1 = errors.New("Сan't divide")

func Err2(x, y int) (int, error) {
	if y == 0 {
		return 0, Err1
	}
	return x / y, nil
}

func main() {
	fmt.Println(Err2(8, 0))

}

// Проверка на наличие ошибки: Напишите функцию, которая возвращает
// ошибку в некоторых случаях. Вызовите эту функцию и проверьте, есть ли
// ошибка, прежде чем продолжать выполнение кода.
