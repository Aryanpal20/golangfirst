
// 53. how to print even numbers of the array

package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		if arr[i]%2 == 0 {
			fmt.Println("the even number of an array", arr[i])
		}
	}
}