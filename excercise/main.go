package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) GetCircumference() float64 {
	return 2 * c.Radius * 3.14
}

func (c Circle) GetArea() float64 {
	return 3.14 * math.Pow(c.Radius, 2)
}

func main() {
	newCircle := Circle{
		Radius: 10,
	}
	fmt.Println(newCircle.GetArea())
	fmt.Println(newCircle.GetCircumference())
}
