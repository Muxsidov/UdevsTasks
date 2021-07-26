package main

import (
	"fmt"
)

// Creating struct
type Contact struct {
	id int
	//phone     string
	firstname string
	//lastname  string
}

type ContactManager struct {
	contacts []Contact
}

func NewContactManager() ContactManager {
	cm := ContactManager{}
	cm.contacts = []Contact{}

	return cm
}

func main() {

	c0 := &Contact{
		id: 0,
		//phone:     "555-00-00",
		firstname: "James",
		//lastname:  "Bond",
	}

	c1 := &Contact{
		id: 1,
		//phone:     "555-00-01",
		firstname: "Peter",
		//lastname:  "Griffin",
	}

	c2 := &Contact{
		id: 2,
		//phone:     "555-00-02",
		firstname: "David",
		//lastname:  "Through",
	}

	//myContact.prepend(c0)
	//myContact.prepend(c1)
	//myContact.prepend(c2)
}

