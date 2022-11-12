//------- 37. The task is to shift all the zeroes appearing in the array to the end of the array. In Golang
// Input: 1 0 7 0 3

// Output: 1 7 3 0 0

package main

import (
	"fmt"
)

func main() {
	var a = [5]int{1, 0, 7, 0, 3}
	var b = 0
	var c = len(a)
	for i := 0; i < c; i++ {
		if a[i] != 0 {
			a[b], a[i] = a[i], a[b]
			b++
		}
	}
	fmt.Println("shifting the zeroes")
	for i := 0; i < c; i++ {
		fmt.Printf("%v ", a[i])
	}
}
