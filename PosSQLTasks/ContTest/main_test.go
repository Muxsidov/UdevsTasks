package main

import (
	"testing"
)

var cm ContactManager

// var id int

func TestNewContactManager(t *testing.T) {
	var err error
	cm, err = NewContactManager()
	if err != nil {
		t.Error("Can not connect a database: ", err)
	}
}
func TestContactManager_Add(t *testing.T) {
	c := Contact{Name: "Rustam", Gender: "male", Phone: 998998149233, Mail: "rustam@mail.ru"}
	err := cm.Add(&c)
	if err != nil {
		t.Error("Can not create a contact: ", err)
	}
}

func TestContactManager_Update(t *testing.T) {
	c := Contact{Name: "Temur", Gender: "male", Phone: 998998149233, Mail: "temur@gmail.com"}

	err := cm.Update(&c)
	if err != nil {
		t.Error("Can not update a contact: ", err)
	}
}

func TestContactManager_ContactList(t *testing.T) {
	err := cm.ContactList()
	if err != nil {
		t.Error("Can not display contacts: ", err)
	}
}

func TestContactManager_Delete(t *testing.T) {
	c := Contact{Name: "Rustam", Gender: "male", Phone: 998998149233, Mail: "rustam@mail.ru"}
	err := cm.Delete(&c)
	if err != nil {
		t.Error("Can not delete contact: ", err)
	}
}
