package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) DoubleSize() (float64, float64) {
	return r.Width * 2, r.Height * 2
}

func main() {
	r := Rectangle{Width: 3.64, Height: 3.89}
	r.DoubleSize()

	fmt.Println(r.DoubleSize())
}

// Методы, изменяющие значения: Добавьте к структуре "Rectangle" метод
// "DoubleSize", который удваивает размеры прямоугольника.
