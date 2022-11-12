//  52. how to compare two arrays using equal to (==) operator

package main

import (
	"fmt"
)

func main() {
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{4, 2, 3}
	arr3 := [3]int{1, 2, 3}

	if arr1 == arr2 {
		fmt.Println("arr1 and arr2 are similar")
	} else {
		fmt.Println("arr1 and arr2 are not similar")
	}
	if arr1 == arr3 {
		fmt.Println("arr1 and arr3 are similar")
	} else {
		fmt.Println("arr1 and arr3 are not similar")
	}
	if arr2 == arr3 {
		fmt.Println("arr2 and arr3 are similar")
	} else {
		fmt.Println("arr2 and arr3 are not similar")
	}

}
