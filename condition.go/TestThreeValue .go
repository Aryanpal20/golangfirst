// 23. we test the three values to find which one is greater
package main

import (
	"fmt"
)

var a = 3
var b = 4
var c = 5

func main() {
	if a > b && a > c {
		fmt.Println("a is greater than b and c")
	} else if b > c && b > a {
		fmt.Println("b is greater than a and c")
	} else {
		fmt.Println("c is greater than a and b")
	}
}
