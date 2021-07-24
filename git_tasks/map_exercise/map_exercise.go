package main

import (
	//"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

var m map[string]int

func WordCount(s string) map[string]int {

	m = make(map[string]int)

	words := strings.Fields(s)
	for _, word := range words {
		if _, w := m[word]; w {
			m[word] += 1
		} else {
			m[word] = 1
		}

	}

	return m
}

func main() {
	wc.Test(WordCount)
}
