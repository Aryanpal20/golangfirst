//   46. create a slice by slicing an array:

package main

import (
	"fmt"
)

func main() {
	var arr = [3]int{3, 4, 5}
	myslice := arr[1:3]

	fmt.Println(myslice)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))
}
