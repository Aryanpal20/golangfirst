// 26. this code will skip the value of 5
package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
