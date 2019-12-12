package main

import "fmt"

func main() {
	//variable declaration
	var a, b = 50, 56
	grades, marks := "B", 90

	/* hello world test program*/
	fmt.Println("helloworld from go language")
	//decisionMaking(a, b, marks, grades)
	loops(a, b, marks, grades)
}

func loops(a, b, marks int, grades string) {
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
	for i, x := range numbers {
		fmt.Println("value of x =", x, "at", i)
	}

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
