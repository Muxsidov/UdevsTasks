package main

import (
	"fmt"
)

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

func (l *linkedList) prepend(n *Contact) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printContact(firstname string) {
	toPrint := l.head
	for l.length != 0 {
		//fmt.Printf("%d ", toPrint.id)
		//fmt.Printf(toPrint.phone)
		//toPrint = toPrint.next
		l.length--	
		if firstname == toPrint.firstname {
			fmt.Printf("%s's phone number is: %s ", firstname, toPrint.phone)
		}
		toPrint = toPrint.next
	}
	//fmt.Printf("\n")
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

  // Put here the name of the person who's number you looking for
	myContact.printContact("Peter")
	/*
		fmt.Println(c0)
		fmt.Println(c1)
		fmt.Println(c2)
	*/
}
