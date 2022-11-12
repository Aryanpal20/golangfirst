// 24. using nested if
package main

import (
	"fmt"
)

var a = 10

func main() {
	if a >= 30 {
		fmt.Println("a is greater than 30")
		if a >= 40 {
			fmt.Println("a is also greater than 40")
		}
	} else {
		fmt.Println("a is less than 30 and 40")
	}
}
