// https://play.golang.org/p/5FUOzjBa-o

package main

import "fmt"

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("anna"))
	fmt.Println(isPalindrome("not a palindrome"))
}
