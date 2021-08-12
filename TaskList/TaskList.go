package main

import (
	"fmt"
)

type Task struct {
	Id                 int
	TaskName          string
	WhomTaskIsGiven    string
	DeadLine           int
	ExplainingOfTask string
}

type TaskManager struct {
	tasks []Task
}

func NewTaskManager() TaskManager {
	cm := TaskManager{}
	cm.tasks = []Task{}

	return cm
}

func (cm *TaskManager) Add(c *Task) *Task {
	cm.tasks = append(cm.tasks, *c)
	id := len(cm.tasks) - 1
	return &cm.tasks[id]
}

func (cm *TaskManager) Update(c *Task) *Task {
	task := &cm.tasks[c.Id]
	task.TaskName = c.TaskName
	contact.WhomTaskIsGiven = c.WhomTaskIsGiven
	contact.DeadLine = c.DeadLine
	contact.ExplainingOTask = c.ExplainingOfTask
	return contact
}

func (cm *TaskManager) Delete(id int) {
	cm.tasks = append(cm.tasks[:id], cm.tasks[id+1:]...)
}

func (cm *TaskManager) TaskList() {
	for _, c := range cm.tasks {
		c.TaskDetail()
	}
}

func (c *Contact) TaskDetail() {
	fmt.Println("**************************")
	fmt.Println("name:", c.TaskName)
	fmt.Println("id:", c.Id)
	fmt.Println("towho:", c.WhomTaskIsGiven)
	fmt.Println("deadline:", c.DeadLine)
	fmt.Println("explaining:", c.ExplainingOfTask)
	fmt.Println()
}

func (c *TaskManager) GetAllTasks() []Task {
	return c.tasks
}

func Menu() {
	fmt.Println("**************************")
	fmt.Println("*          Menu          *")
	fmt.Println("**************************")

	fmt.Println("List of all tasks - 1")
	fmt.Println("Add new task      - 2")
	fmt.Println("Update a task     - 3")
	fmt.Println("Delete a task     - 4")
	fmt.Println("Exit                 - 5")
	fmt.Println("**************************")
}

func EnterDetails(c *Task) {
	var deadline int
	var name, towho, explaining string

	fmt.Print("Enter name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter to whom: ")
	fmt.Scanln(&towho)
	fmt.Print("Enter deadline: ")
	fmt.Scanln(&deadline)
	fmt.Print("Enter explaining: ")
	fmt.Scanln(&explaining)

	c.Task_Name = name
	c.Whomtaskisgiven = towho
	c.Deadline = deadline
	c.Explaining_of_task = explaining
}

func main() {
	var choice int
	var id int
	var cm TaskManager
	var c Task

	for {
		Menu()
		fmt.Print("Enter a your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			cm.TaskList()
		} else if choice == 2 {
			c.Id = len(cm.tasks)
			EnterDetails(&c)
			cm.Add(&c)
		} else if choice == 3 {
			fmt.Print("Enter id: ")
			fmt.Scanln(&id)

			if id <= len(cm.contacts)-1 {
				c.Id = id
				EnterDetails(&c)
				cm.Update(&c)
			} else {
				fmt.Println("Entered id does not exists.")
			}
		} else if choice == 4 {
			fmt.Print("Enter id: ")
			fmt.Scanln(&id)

			if id <= len(cm.contacts)-1 {
				cm.Delete(id)
			} else {
				fmt.Println("Entered id does not exists.")
			}
		} else if choice == 5 {
			break
		}
	}
}
