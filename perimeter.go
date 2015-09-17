package main

import (
	"fmt"
	"math"
)


type Shaper interface {
	Area() float64
	Perimeter() float64
}



type Rect struct {
	length, width float64
}

func (r Rect) Area() float64 {
	return r.length * r.width
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.length + r.width)
}



type Circle struct {
	r float64
}


func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.r
}



func main() {
	r := Rect{length: 5, width: 3}
	fmt.Println(" The Details of rectangle are ", r)

	var x Shaper
	x = r 
	fmt.Println("Shaper Area(Rect) ", x.Area())
	fmt.Println()

	fmt.Println("Shaper Perimeter(Rect): ", x.Perimeter())
	fmt.Println()

	c := Circle{r: 5}
	fmt.Println("The Details of circle are: ", c)
	x = c 
	fmt.Println("Shaper Area(Circle): ", x.Area())
	fmt.Println()

	fmt.Println("Shaper Perimeter(Circle): ", x.Perimeter())
	fmt.Println()
}