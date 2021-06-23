package main

import (
	"fmt"
)

func main() {
	
	evensum, oddsum := 0, 0
	end := 10 // How many numbers you want ?

	for i := 1; i < end; i++ {

		if i%2 == 0 {
			evensum += i
		} else {
			oddsum += i
		} 
	}
	fmt.Printf("evensum: %v \noddsum: %v\n", evensum, oddsum)
}

