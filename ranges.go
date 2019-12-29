package main

import "fmt"

/*RangesEx is a test function for ranges in Go*/
func RangesEx() {
	/*delcering a map*/
	marksMap := make(map[string]int)

	/*initailizing values to the map*/
	marksMap["dhamu"] = 79
	marksMap["meena"] = 90

	/*Printing the values stored in the map*/
	for name, marks := range marksMap {
		fmt.Printf("The mark of %s is : %d\n", name, marks)
	}
}
