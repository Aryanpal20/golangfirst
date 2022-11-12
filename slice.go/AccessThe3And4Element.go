// 48.  to access the first and third elements in the slice:

package main

import (
	"fmt"
)

func main() {
	myslice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(myslice[0])
	fmt.Println(myslice[2])
}
