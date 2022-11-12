// 21. we test the two values to find out which one is greater
package main

import (
	"fmt"
)

var a = 4
var b = 5

func main() {
	if a > b {
		fmt.Println("a is greater than b")
	} else {
		fmt.Println("b is greater than a")
	}
}
