// 29. second largest element in array
package main

import (
	"fmt"
)

func main() {
	var arr = [4]int{1, 5, 3, 4}
	var max1 = 0
	var max2 = 0

	for i := 0; i <= 3; i++ {
		if max1 < arr[i] {
			max2 = max1
			max1 = arr[i]
		} else if max2 < arr[i] {
			max2 = arr[i]
		}
	}
	fmt.Println(max2)
}
