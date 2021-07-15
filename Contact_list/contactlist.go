package main

import (
	"fmt"
)

// Creating struct 
type Contact struct {
	id        int
	phone     string
	firstname string
	lastname  string
	next      *Contact
}

type linkedList struct {
	head   *Contact
	length int
}

// Adding to struct 
func (l *linkedList) prepend(n *Contact) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

// Searching for phone number by name
func (l linkedList) printContact(firstname string) {
	toPrint := l.head
	for i := 0; i < l.length; i++ {
		//l.length--
		if firstname == toPrint.firstname {
			fmt.Printf("%s's phone number is: %s ", firstname, toPrint.phone)
			return
		} 
		toPrint = toPrint.next
	}
}

// Deleting ...
func (l *linkedList) deleteContact(firstname string) {
	if l.length == 0 {
		return
	}

	if l.head.firstname == firstname {
		l.head = l.head.next
		l.length--
		return
	}
	
	previousToDelete := l.head
	for previousToDelete.next.firstname != firstname {
		if previousToDelete.next.next ==nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
		
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	// 	fmt.Println("Hello, playground")

	myContact := linkedList{}

	c0 := &Contact{
		id:        0,
		phone:     "555-00-00",
		firstname: "James",
		lastname:  "Bond",
	}

	c1 := &Contact{
		id:        1,
		phone:     "555-00-01",
		firstname: "Peter",
		lastname:  "Griffin",
	}

	c2 := &Contact{
		id:        2,
		phone:     "555-00-02",
		firstname: "David",
		lastname:  "Through",
	}

	myContact.prepend(c0)
	myContact.prepend(c1)
	myContact.prepend(c2)
	
	// Put person's name to delete
	myContact.deleteContact("David")
	
	// Put person's name whom number you want to see
	myContact.printContact("Peter")
}
