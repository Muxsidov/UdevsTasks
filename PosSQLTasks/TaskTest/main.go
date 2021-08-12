// https://www.calhoun.io/using-postgresql-with-go/

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Collected information abotu database
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234qwer"
	dbname   = "Contactlistdb"
)

type TaskManager struct {
	db *sql.DB
}

// Create a struct for SELECT *
type Task struct {
	Id              int
	TaskName        string
	WhomTaskIsGiven string
	Deadline        int
	Explaining      string
}

func NewTaskManager() (TaskManager, error) {
	cm := TaskManager{}
	var err error

	// Creating the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Opening and checking if we opened connector correctly
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return cm, nil
}

func Menu() {
	fmt.Println("**************************")
	fmt.Println("*          Menu          *")
	fmt.Println("**************************")

	fmt.Println("List of tasks     - 1")
	fmt.Println("Add new task      - 2")
	fmt.Println("Update a task     - 3")
	fmt.Println("Delete a task     - 4")
	fmt.Println("Exit              - 5")
	fmt.Println("**************************")
}

func (cm *TaskManager) Add(c *Task) error {
	// func Add(c *Contact) error{

	// Inserting
	sqlStatement := `
	INSERT INTO tasks (name, towho, deadline, explaining)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 0
	cm.db.QueryRow(sqlStatement, &c.TaskName, &c.WhomTaskIsGiven, &c.Deadline, &c.Explaining).Scan(&id)
	fmt.Println("New record ID is:", id)
	return nil
}

func (cm *TaskManager) Update(c *Task) {
	var err error

	// Updating
	sqlStatement := `
	UPDATE tasks
	SET name = $2, towho = $3, deadline = $4, explaining = $5
	WHERE id = $1
	RETURNING id, name;`
	var name string
	var id int
	err = cm.db.QueryRow(sqlStatement, &c.TaskName, &c.WhomTaskIsGiven, &c.Deadline, &c.Explaining).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println(id, name)
}

func (cm *TaskManager) Delete(c *Task) {
	// Deleting
	sqlStatement := `
	DELETE FROM tasks
	WHERE id = $1;`
	res, err := cm.db.Exec(sqlStatement, &c.Id)
	if err != nil {
		panic(err)
	}

	// Cheking if Updating/Deleting was succsessful
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)

}

func (cm *TaskManager) TaskList() {

	rows, err := cm.db.Query("SELECT id, name, deadline, explaining FROM contacts LIMIT $1", 3)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, deadline int
		var name, towho string
		err = rows.Scan(&id, &name, &deadline, &towho)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(id, name, deadline, towho)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}

func EnterDetails(c *Task) {
	var deadline int
	var name, explaining, towho string

	fmt.Print("Enter name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter to who: ")
	fmt.Scanln(&towho)
	fmt.Print("Enter deadline: ")
	fmt.Scanln(&deadline)
	fmt.Print("Enter explaining: ")
	fmt.Scanln(&explaining)

	c.TaskName = name
	c.WhomTaskIsGiven = towho
	c.Deadline = deadline
	c.Explaining = explaining
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
			EnterDetails(&c)
			cm.Add(&c)
		} else if choice == 3 {
			fmt.Print("Enter id: ")
			fmt.Scanln(&id)
			c.Id = id
			EnterDetails(&c)
			cm.Update(&c)
		} else if choice == 4 {
			fmt.Print("Enter id: ")
			fmt.Scanln(&id)
			c.Id = id
			cm.Delete(&c)
		} else if choice == 5 {
			break
		}
	}
}
