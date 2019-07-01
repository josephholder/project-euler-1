package main

import (
	"fmt"
)

func main() {

	value := 0

	three := 3
	five := 5

	for i := 0; i < 1000; i++ {
		if i%three == 0 || i%five == 0 {
			value += i
		}
	}

	fmt.Println(value)
}
