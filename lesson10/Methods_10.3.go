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

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	r := Rectangle{Width: 3.64, Height: 3.89}
	r.IsSquare()
	r.Area()

	fmt.Println("Is this rectangle a square?", r.IsSquare(), "The area of the figure is ", r.Area())

	t := Rectangle{Width: 8.66, Height: 8.66}
	t.IsSquare()
	t.Area()
	fmt.Println("Is this rectangle a square?", t.IsSquare(), "The area of the figure is ", t.Area())
}

// 	Использование нескольких методов: Используйте методы "Area" и
// "IsSquare" на нескольких экземплярах структуры "Rectangle".
