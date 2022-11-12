// 19. initialize a specific element of an array
package main

import (
	"fmt"
)

func main() {
	var arr1 = [4]int{1: 20, 2: 30}
	fmt.Println(arr1)
	fmt.Println("")
	// initialize a specific element of an array by give them
	var arr2 = [4]int{1: 20, 2: 30}
	arr2[0] = 10
	arr2[3] = 40
	fmt.Println(arr2)
}
