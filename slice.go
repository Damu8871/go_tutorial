package main

import "fmt"

/*SliceBasic is a test progrmam for slice in Go*/
func SliceBasic() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sli := array[2:4]
	sli1 := array[1:7]
	fmt.Println("sli == nil", sli == nil, sli)
	SliceComplete(sli)
	SliceComplete(sli1)
}

/*SliceComplete is a test program function for slice in go*/
func SliceComplete(slicearr []int) {
	fmt.Printf("Length of Slice is :%d\n Capacity of Slice is :%d\n The Sliced values are :%v\n", len(slicearr), cap(slicearr), slicearr)
}
