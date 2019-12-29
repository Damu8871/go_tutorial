package main

import (
	"fmt"
	"math"
)

func main() {
	//variable declaration
	// var a, b = 50, 56
	// grades, marks := "B", 90
	// pulse, sep := 23, 34

	/* hello world test program*/
	fmt.Println("helloworld from go language")

	/*Strings in go
	Strings()*/

	/*StringArray in go
	StringArray()*/

	/*Pointers in Go
	PointersEx()*/

	/*Array of pointers in Go
	ArrayofPointers()*/

	/*Pointer of pointer in Go
	PointerofPointer()*/

	/*Pass pointers as parameters
	PassPointers()*/

	/*Structures in Go
	StructTest()*/

	/*SliceBasic in Go
	SliceBasic()
	SliceCopy()*/

	/*Ranges in Go*/
	//RangesEx()
	RangesEx1()

	/*Go Methods
	circle := Circle{x: 0, y: 0, radious: 5}
	fmt.Println("Radious of the circle is :", circle.area())*/

	/*Function Closure
	nextNumber := getSequence()
	nextNumber1 := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	for m := 0; m <= 3; m++ {
		fmt.Println(nextNumber1())
	}*/

	/*Function as value
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9)) */

	/*Decision making in Go
	decisionMaking(a, b, marks, grades)*/

	/* CallbyValue in Go
	callByValueReturns := loops(a, b, pulse, sep, marks, grades)
	fmt.Println("loop returns :", callByValueReturns)*/

	/*Call by Reference
	c1, c2 := callByReferenceEx(&a, &b, pulse, sep, marks, grades)
	fmt.Println("Call by reference returns: ", c1, c2, "after a and b :", a, b)*/

}

/*Circle is just a type contains variables */
type Circle struct {
	x, y, radious float64
}

/*Create a method*/
func (circle Circle) area() float64 {
	return math.Pi * circle.radious * circle.radious
}

func getSequence() func() int {
	i := 0
	return func() int {
		i = i + 1
		return i
	}
}

func callByReferenceEx(a *int, b *int, pulse, sep, mark int, grades string) (int, int) {
	var temp int
	temp = *a /* save the value at address x */
	*a = *b   /* put y into x */
	*b = temp /* put temp into y */
	return *a, *b
}

func loops(a, b, pulse, sep, marks int, grades string) int {
	//variable declaration
	numbers := [6]int{2, 4, 5, 6}
	// loops of execution
	for val := 3; val <= 10; val++ {
		fmt.Println("value of val:", val)
	}
	for a < b {
		a++
		fmt.Println("a is :", a)
	}
	//range of numbers in loop
	for i, x := range numbers {
		fmt.Println("value of x =", x, "at", i)
	}
	//Nested Loop
	j := 2
	for p := 2; p < 100; p++ {
		for j := 2; j <= (p / j); j++ {
			if (p % j) == 0 {
				break
			}
		}
		if j > (p / j) {
			fmt.Println("The prime number is", p)
		}
	}
	for ; pulse <= 40; pulse++ {
		fmt.Println("Pulse:", pulse)
		for ; sep >= pulse; pulse++ {
			fmt.Println("Pulse:", pulse)
		}
	}

	//Reinitializing pulse and declaring FOR loop
	//BREAK Statement
	for pulse := 23; pulse <= sep; pulse++ {
		fmt.Println("Pulse from break statement :", pulse)

		if pulse >= 30 {
			break
		}
	}

	//Continue Statement
	for pulse := 30; pulse <= 40; {
		if pulse == 35 || pulse == 37 {

			fmt.Println("Skipping pulse using continue:", pulse)
			pulse++
			continue
		}
		fmt.Println("Pulse form continue:", pulse)
		pulse++
	}

	//Goto Statement
	//nov var initialization
	nov := 10
Start:
	for nov <= 20 {
		if nov == 15 || nov == 19 {
			fmt.Println("Skipping from goto statement:", nov)
			nov++
			goto Start
		}
		fmt.Println("Nov from goto statement:", nov)
		nov++
	}
	return pulse
}

func decisionMaking(num1, num2, marks int, grades string) {
	fmt.Println("if statement")
	if num1 > num2 {
		/* if condition is true then print the following */
		fmt.Println("A is greater than B")
	} else if num1 == num2 {
		// if above condition is false and this condition is true
		fmt.Println("Both A and B are equal")
	} else {
		/* if and else if condition is false then print the following */
		fmt.Println("else if failed - else block")
	}
	if num1 == 34 {
		if num2 == 56 {
			//Nested if condition
			fmt.Println("Nested if condition passed")
		}
	} else {
		fmt.Println("Initial Nested if fails")
	}

	//Switch statement
	switch marks {
	case 90:
		grades = "A"
	case 80:
		grades = "B"
	case 50, 60, 70:
		grades = "C"
	default:
		grades = "D"
	}
	switch {
	case grades == "A":
		fmt.Println("Excellent!! ")
	case grades == "B", grades == "C":
		fmt.Println("Well Done! ")
	case grades == "D":
		fmt.Println("Just Passed ")
	}
}
