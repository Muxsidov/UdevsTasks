// https://tutorialedge.net/golang/intro-testing-in-go/

package main

import (
	"fmt"
	"testing"
)

// Calculate returns x + 2.
func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func TestCalculate(t *testing.T) {
    if Calculate(2) != 4 {
        t.Error("Expected 2 + 2 to equal 4")
    }
}

func TestTableCalculate(t *testing.T) {
    var tests = []struct {
        input    int
        expected int
    }{
        {2, 4},
        {-1, 1},
        {0, 2},
        {-5, -3},
        {99999, 100001},
    }

    for _, test := range tests {
        if output := Calculate(test.input); output != test.expected {
            t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
        }
    }
}

func main() {
	fmt.Println("Hello World")
}
