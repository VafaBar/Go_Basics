package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) IsSquare() bool {
	if r.Width == r.Height {
		return true
	} else {
		return false
	}
}

func main() {
	r := Rectangle{Width: 3.64, Height: 3.89}
	r.IsSquare()

	fmt.Println("Is this rectangle a square?", r.IsSquare())
}

//  Методы с разными возвращаемыми значениями: Добавьте к структуре
// "Rectangle" метод "IsSquare", который возвращает bool, указывающий,
// является ли прямоугольник квадратом.
