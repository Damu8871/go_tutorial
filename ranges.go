package main

import "fmt"

/*RangesEx is a test function for ranges in Go*/
func RangesEx() {
	/*delcering a map*/
	marksMap := make(map[string]int)

	/*initailizing values to the map*/
	marksMap["mark"] = 79
	marksMap["spencer"] = 90

	/*Printing the values stored in the map*/
	for name, marks := range marksMap {
		fmt.Printf("The mark of %s is : %d\n", name, marks)
	}
}

/*RangesEx1 is an another example of ranges in go*/
func RangesEx1() {
	/*Array example*/
	arr := []int{23, 4, 56, 12}
	fmt.Println("Array Example")
	for ind, val := range arr {
		fmt.Printf("The index of %x is %d\n", ind, val)
	}

	/*String example*/
	name := "Micheal"
	fmt.Println("String Example")
	for ind, code := range name {
		fmt.Printf("The index of %c is %d\n", code, ind)
	}

	/*Country example*/
	fmt.Println("Country Example")
	/*Creating map variable*/
	countryCapitalMap := map[string]string{"India": "Delhi", "France": "Paris", "Japan": "Tokyo"}

	for country := range countryCapitalMap {
		fmt.Printf("The capital of %s is %s\n", country, countryCapitalMap[country])
	}
}
