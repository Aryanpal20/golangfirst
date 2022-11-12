// 51.  how to calculate the sum of all array elements

package main

import (
	"fmt"
)

func main() {
	arr := [5]int{2, 3, 4, 5, 6}
	sum := 0
	for i := 0; i < 5; i++ {
		sum += arr[i]
	}
	fmt.Println("the sum of array is :- ", sum)
}
