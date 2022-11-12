//              // 22. we test the two values to find greater or equal
package main

import (
	"fmt"
)

var a = 6
var b = 6

func main() {
	if a > b {

		fmt.Println("a is greater than b")
	} else if b > a {

		fmt.Println("b is greater than a")
	} else {
		fmt.Println("both are equal ")
	}
}
