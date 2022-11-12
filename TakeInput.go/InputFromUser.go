// 28. write a code to take inout from user
package main

import (
	"fmt"
)

func main() {

	var first string
	fmt.Println("enter your first name :- ", first)
	fmt.Scanln(&first)

	var last string
	fmt.Println("enter your last name :- ", last)
	fmt.Scanln(&last)
	fmt.Println("your full name is :-")
	fmt.Println(first + " " + last)
}
