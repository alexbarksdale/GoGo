package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}
type Rect struct {
	Width, Height float64
}
type Cuboid struct {
	Rect
	Length float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rect) Area() float64 {
	return r.Width * r.Height
}

func (q Cuboid) Area() float64 {
	return 2 * (q.Width*q.Height + q.Width*q.Length + q.Height*q.Length)
}

func measurePrint(s Shape) {
	fmt.Printf("%T %+v\n", s, s)
	fmt.Printf("Area: %v\n\n", s.Area())

}

func main() {
	c := Circle{Radius: 5}
	r := Rect{Width: 2, Height: 5}
	q := Cuboid{Rect: r, Length: 5}

	measurePrint(c)
	measurePrint(r)
	measurePrint(q)
}
