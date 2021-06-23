package main

import "fmt"

func main() {
	first, second := 0, 1

	end := 15 // How many numbers do you want ?

	for i := 0; i <= end-1; i++ {
		fmt.Println(first)
		temp := first + second
		first = second
		second = temp
	}
}
