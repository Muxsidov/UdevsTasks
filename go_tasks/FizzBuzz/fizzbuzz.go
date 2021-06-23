package main

import (
	"fmt"
)

func main() {

	end := 16 // How many numbers you want ?

	for i := 1; i < end; i++ {
		//fmt.Println("Hello, playground")

		if i%5 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}
