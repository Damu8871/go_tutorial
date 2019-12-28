package main

import "fmt"

/*SliceBasic is a test progrmam for slice in Go*/
func SliceBasic() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sli := array[2:4]
	sli1 := array[1:7]
	fmt.Println("sli == nil", sli == nil, sli)
	/*Calling SliceComplete function from this function*/
	SliceComplete(sli)
	SliceComplete(sli1)

	/*Calling slice append function from this function*/
	SliceAppend(array)
}

/*SliceComplete is a test program function for slice in go*/
func SliceComplete(slicearr []int) {
	fmt.Printf("Length of Slice is :%d\n Capacity of Slice is :%d\n The Sliced values are :%v\n", len(slicearr), cap(slicearr), slicearr)
}

/*SliceAppend is a test program for append in slice*/
func SliceAppend(arr []int) {
	sli := arr[3:8]
	fmt.Printf("before -> array = %v\n", arr)
	fmt.Printf("before -> sli = %v\n", sli)
	fmt.Printf("before -> len = %d cap = %d\n", len(sli), cap(sli))
	fmt.Println("&arr[3] == &sli[0]", &arr[3], &sli[0], &arr[3] == &sli[0])

	/*Appending an array to the sliced array*/
	sli = append(sli, 10, 20, 30, 40)
	fmt.Printf("after -> array = %v\n", arr)
	fmt.Printf("after -> sli = %v\n", sli)
	fmt.Printf("after -> len = %d cap = %d\n", len(sli), cap(sli))
	fmt.Println("&arr[3] == &sli[0]", &arr[3], &sli[0], &arr[3] == &sli[0])
}

/*SliceCopy is a program for copying array to another array*/
func SliceCopy() {
	s1 := make([]int, 5, 5)
	s2 := []int{1, 2, 3, 4}

	fmt.Printf("before -> s1 = %v\n", s1)
	fmt.Printf("before -> s2 = %v\n", s2)

	n1 := copy(s1, s2)
	fmt.Printf("after -> s1 = %v\n", s1)
	fmt.Printf("after -> s2 = %v\n", s2)
	fmt.Printf("n1 = %d\n", n1)
}
