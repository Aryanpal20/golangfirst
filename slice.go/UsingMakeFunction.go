// 47. The make() function can also be used to create a slice.

package main

import (
	"fmt"
)

func main() {
	myslice := make([]int, 4)
	myslice = append(myslice, 1, 2, 3, 4)
	myslice = append(myslice[4:]) // we can use this code for remove first 4 zero
	fmt.Println(myslice)
}
