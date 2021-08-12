package main

import (
	"testing"
)

var cm TaskManager

// var id int

func TestNewTaskManager(t *testing.T) {
	var err error
	cm, err = NewTaskManager()
	if err != nil {
		t.Error("Can not connect a database: ", err)
	}
}
func TestNewTaskManager_Add(t *testing.T) {
	c := Task{TaskName: "ContactList", WhomTaskIsGiven: "Rustam", Deadline: 15082021, Explaining: "CRUD"}
	err := cm.Add(&c)
	if err != nil {
		t.Error("Can not create a contact: ", err)
	}
}

func TestTaskManager_Update(t *testing.T) {
	c := Task{TaskName: "TaskList", WhomTaskIsGiven: "Temur", Deadline: 1082021, Explaining: "TaskList"}

	err := cm.Update(&c)
	if err != nil {
		t.Error("Can not update a contact: ", err)
	}
}

func TestTaskManager_TaskList(t *testing.T) {
	err := cm.TaskList()
	if err != nil {
		t.Error("Can not display contacts: ", err)
	}
}

func TestTaskManager_Delete(t *testing.T) {
	c := Task{TaskName: "ContactList", WhomTaskIsGiven: "Rustam", Deadline: 15082021, Explaining: "CRUD"}
	err := cm.Delete(&c)
	if err != nil {
		t.Error("Can not delete contact: ", err)
	}
}
