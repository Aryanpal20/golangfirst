// 49. change a specific slice element by referring to the index number.

package main

import (
	"fmt"
)

func main() {
	myslice := []int{1, 2, 3, 4, 5}
	myslice[3] = 5
	fmt.Println(myslice)
}
