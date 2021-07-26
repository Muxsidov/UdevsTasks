package main

// importing required packages
import (
	"fmt"
)

type id struct {
	id int
}

type task_list struct {
	name     string
	deadline string
}



func main() {
	id_0 := id{0}
	task1 := task_list{name: "Declare an array", deadline: "01.22.2020"}

	//var mp map[id]task_list

	sample := map[id]task_list{id_0: task1}
	fmt.Print(sample)

}

