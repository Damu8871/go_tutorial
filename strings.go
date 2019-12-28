package main

import (
	"fmt"
	"strings"
)

/*Strings is a function for learning*/
func Strings() {
	var greetings = "Hello world!!"

	/*printing string as words*/
	fmt.Printf("Normal String:")
	fmt.Printf("%s", greetings)
	fmt.Printf("\n")
	fmt.Printf("hex bytes: ")

	for i := 0; i < len(greetings); i++ {
		fmt.Printf("%x ", greetings[i])
	}
	fmt.Printf("\n")

	/*Join function in strings*/
	strJoin := []string{"My", "Code"}
	fmt.Println(strings.Join(strJoin, " "))
}

/*StringArray is a function to learn about it*/
func StringArray() {
	emptyArray := [10]int{}
	//var emptyArray [10]int

	for i := 0; i < 10; i++ {
		emptyArray[i] = i + 100
		fmt.Printf("Value in emptyArray[%d] = %d\n", i, emptyArray[i])
	}
	averageVal := CalculateAverage(emptyArray)
	fmt.Println("average value : ", averageVal)
}

/*CalculateAverage is a function to calculate the average of numbers in an array*/
func CalculateAverage(averageArray [10]int) int {
	var total int
	for i := 0; i < len(averageArray); i++ {
		total += averageArray[i]
	}
	average := total / len(averageArray)
	return average
}
