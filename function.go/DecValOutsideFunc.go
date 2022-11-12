// 5. this will show declaring variable outside of  a function
//  we can use only var keyword outside the function
package main

import (
	"fmt"
)

var x int
var y int = 2
var z = 3

func main() {
	x = 1
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
