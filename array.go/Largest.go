// 30. largest element in array
package main

import (
	"fmt"
)

func main() {
	var arr = [4]int{1, 5, 3, 4}
	var max = arr[0]
	for i := 0; i <= 3; i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	fmt.Println((max))
}
