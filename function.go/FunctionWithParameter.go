// 42. function with parameter

package main

import (
	"fmt"
)

func information(name string, mark int, subject string) {
	fmt.Println("my name is ", name)
	fmt.Println("my mark is ", mark)
	fmt.Println("the subject is ", subject)
}
func main() {
	information("rohan", 78, "maths")
	information("mohan", 80, "maths")
}
