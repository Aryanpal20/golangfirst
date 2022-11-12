// 27. this code will break the value after 5
package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}
