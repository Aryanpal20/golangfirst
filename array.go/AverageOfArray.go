// 38.  -----Given an array of n elements, your task is to find out the average of the array.
//---- Approach:

//------------- Accept the size of the array.
// ---------Accept the elements of the array.
// ---------Store the sum of the elements using for loop.
// ----------Calculate Average = (sum/size of array)
// ----------Print the average.\

package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4}
	sarr := 4
	s := 0
	for i := 0; i < sarr; i++ {

		s = s + arr[i]
	}
	avg := (float64(s)) / (float64(sarr))
	fmt.Println(s)
	fmt.Println(avg)
}
