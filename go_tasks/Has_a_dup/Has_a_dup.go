package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 10}

	for i := 0; i < len(arr); i++ {
		// fmt.Println(arr[i])
		for j := i + 1; j < len(arr); j++ {
			if arr[j] == arr[i] {
				fmt.Println(arr[i])
			}
		}
	}
}
