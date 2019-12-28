package main

import "fmt"

/*PointersEx is a function to know about pointers in Go*/
func PointersEx() {
	var val int = 20
	var add *int
	add = &val
	fmt.Printf("The val is :%d\n", val)
	fmt.Printf("The address of val is :%x\n", add)
	fmt.Printf("The value in add is :%d\n", *add)
	/*Nil pointer in Go*/
	var ptr *int
	fmt.Printf("The address saved in the ptr is :%x\n", ptr)
}

/*ArrayofPointers is an example function for arroy of pointers*/
func ArrayofPointers() {
	var arr = []int{10, 20, 30}
	var ptr [3]*int

	for i := 0; i < len(arr); i++ {
		ptr[i] = &arr[i]
		fmt.Printf("The arr is : %d and the address is : %x\n", arr[i], ptr[i])
	}
}

/*PointerofPointer is an example of Pointer to Pointer*/
func PointerofPointer() {
	var val = 300
	var ptr *int
	var pptr **int

	ptr = &val  /*assinging address of val to ptr*/
	pptr = &ptr /*assinging address of ptr to pptr*/

	fmt.Printf("the value is : %d\n", val)
	fmt.Printf("value in address of %d is : %d\n", val, *ptr)
	fmt.Printf("address of value in %d stored in ptr is : %d\n", *ptr, **pptr)
}

/*PassPointers is an exmaple of Passing pointer to function*/
func PassPointers() {
	var num = 4
	var num1 = 5
	PointerAsParameter(&num, &num1)
	fmt.Printf("num and num1 are : %d, %d\n", num, num1)
}

/*PointerAsParameter is an example function calling from PassPointers function*/
func PointerAsParameter(ptr *int, ptr1 *int) {
	temp := *ptr
	*ptr = *ptr1
	*ptr1 = temp
}
