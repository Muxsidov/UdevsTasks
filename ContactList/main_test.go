package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	arg := Contact {
		Id: 1,
		Name: "Peter",
		Gender: "male",
		Phone: 5550002,
		Mail: "pdgr@gmail.com",
	}
	
	Add(arg)
	
	if contains(s, Peter) != true {
		t.Error("There is no such name")
	}
	
}

func contains(s []string, str string)bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	
	return false
}
