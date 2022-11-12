//  55. how to reverse an integer array

package main

import (
	"fmt"
)

func main() {
	var arr1 [4]int
	var j int = 3
	arr := [4]int{1, 2, 3, 4}
	for i := 0; i < 4; i++ {
		arr1[j] = arr[i]
		j = j - 1
	}
	for i := 0; i < 4; i++ {
		fmt.Printf("%d ", arr1[i])
	}
}
