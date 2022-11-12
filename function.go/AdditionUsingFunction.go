// 43. myFunction() receives two integers (x and y) and returns their addition (x + y) as integer (int)

package main

import (
	"fmt"
)

func integer(x int, y int) (result int) {
	result = x + y
	return result
}
func main() {
	fmt.Println(integer(2, 4))
	fmt.Println(integer(10, 20))
}
