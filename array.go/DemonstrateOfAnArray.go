//  39. Golang program to demonstrate the example of an Array

package main

import (
	"fmt"
)

func main() {
	var arr = [4]int{1, 2, 3, 4}
	fmt.Println("the array is :-")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
}
