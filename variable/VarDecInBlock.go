//  9. variable declartion in a block
package main

import (
	"fmt"
)

func main() {
	var (
		a, b, c = 1, 2, 3
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
