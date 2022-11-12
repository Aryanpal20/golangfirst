// 44. we store the return value in a variable

package main

import (
	"fmt"
)

func valueinvariable(x int, y int) (result int) {
	result = x + y
	return
}
func main() {
	var total int
	total = valueinvariable(2, 6)
	fmt.Println(total)
}
