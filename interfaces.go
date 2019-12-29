package main

import "fmt"

/*Shapes is an interface */
type Shapes interface {
	area() float64
	perimeter() float64
}

/*Rect is a struct*/
type Rect struct {
	width  float64
	height float64
}

/*area is a fuction with interfaces*/
func (r Rect) area() float64 {
	return r.width * r.height
}

/*perimeter is a fuction with interfaces*/
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

/*InterfacesEx is a test function for interfaces in go*/
func InterfacesEx() {
	var s Shapes
	s = Rect{2.5, 7.2}
	r := Rect{3.0, 9.6}
	fmt.Println("The area of the rectangle is :", s.area())
	fmt.Println("The perimeter of the rectange is :", s.perimeter())
	fmt.Printf("The area of the rectangle is :%0.2f\n", r.area())
	fmt.Println("The perimeter of the rectange is :", r.perimeter())
}
