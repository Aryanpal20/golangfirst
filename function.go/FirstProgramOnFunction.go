//  42. first program on function

package main

import (
	"fmt"
)

func myMessage() {
	fmt.Println("this is my first function program")
}
func main() {
	myMessage()
	// here we can call many times
	myMessage()
	myMessage()
	myMessage()
}
