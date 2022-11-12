// 50. how to append one slice to another slice:

package main

import (
	"fmt"
)

func main() {
	myslice1 := []int{1, 2, 3}
	myslice2 := []int{4, 5, 6}
	myslice3 := append(myslice1, myslice2...)
	myslice4 := []int{7, 8, 9}
	myslice5 := append(myslice3, myslice4...)
	fmt.Println(myslice5)
}
