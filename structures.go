package main

import "fmt"

/*Book is a structure variable*/
type Book struct {
	title  string
	author string
	bookid int
}

/*StructTest is a function to test the variables set in structure variable*/
func StructTest() {
	/*Decalring a variable from structure variable*/
	var book1 Book

	/*Assinging values to the variables created just above*/
	book1.title = "Go programming"
	book1.author = "Google"
	book1.bookid = 765842

	/*Printing the values in the variables*/
	fmt.Printf("Title of Book1 is : %s\n", book1.title)
	fmt.Printf("Author of Book1 is : %s\n", book1.author)
	fmt.Printf("BookID of Book1 is : %d\n", book1.bookid)

	/*Passing the structure variable as function arguments*/
	PassingStruct(&book1)
}

/*PassingStruct used to pass the structure variable as an argument*/
func PassingStruct(bookOne *Book) {
	/*Printing values from the arguments passsed by*/
	fmt.Printf("The value in struct variabels is :%s\n", bookOne.author)
	fmt.Printf("The value in struct variabels is :%s\n", bookOne.title)
	fmt.Printf("The value in struct variabels is :%d\n", bookOne.bookid)
}
