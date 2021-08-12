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

// Create a struct for SELECT *
type Task struct {
	Id     int
	Task_Name   string
	Whomtaskisgiven string
	Deadline  int
	Explanation   string
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

func Add(c *Task) {
	// Creating the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Opening and checking if we opened connector correctly
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Inserting
	sqlStatement := `
	INSERT INTO contacts (task_name, who, deadline, explanation)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, &c.Task_Name, &c.Whomtaskisgiven, &c.Deadline, &c.Explanation).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}

func Update(c *Task) {

	// Creating the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Opening and checking if we opened connector correctly
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Updating
	sqlStatement := `
	UPDATE contacts
	SET task_name = $2, who = $3, deadline = $4, explanation = $5
	WHERE id = $1
	RETURNING id, name;`
	var name string
	var id int
	err = db.QueryRow(sqlStatement, &c.Id, &c.Task_Name, &c.Whomtaskisgiven, &c.Deadline, &c.Explanation).Scan(&id, &name)
	if err != nil {
		panic(err)
	}
	fmt.Println(id, name)
}

func Delete(c *Task) {

	// Creating the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Opening and checking if we opened connector correctly
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Deleting
	sqlStatement := `
	DELETE FROM contacts
	WHERE id = $1;`
	res, err := db.Exec(sqlStatement, &c.Id)
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

func TaskList() {

	// Creating the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Opening and checking if we opened connector correctly
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, task_name, deadline, explanation FROM contacts LIMIT $1", 3)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, deadline int
		var task_name, explanation string
		err = rows.Scan(&id, &task_name, &deadline, &explanation)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(id, task_name, deadline, explanation)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}

func EnterDetails(c *Task) {
	var deadline int
	var task_name, who, explanation string

	fmt.Print("Enter task_name: ")
	fmt.Scanln(&task_name)
	fmt.Print("Enter whom task was given: ")
	fmt.Scanln(&who)
	fmt.Print("Enter deadline: ")
	fmt.Scanln(&deadline)
	fmt.Print("Enter explanation: ")
	fmt.Scanln(&explanation)

	c.Task_Name = task_name
	c.Whomtaskisgiven = who
	c.Deadline = deadline
	c.Explanation = explanation
}

func main() {

	var choice int
	var id int
	var c Task

	for {
		Menu()
		fmt.Print("Enter a your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			TaskList()
		} else if choice == 2 {
			EnterDetails(&c)
			Add(&c)
		} else if choice == 3 {
			fmt.Print("Enter id: ")
			fmt.Scanln(&id)
			c.Id = id
			EnterDetails(&c)
			Update(&c)
		} else if choice == 4 {
			fmt.Print("Enter id: ")
			fmt.Scanln(&id)
			c.Id = id
			Delete(&c)
		} else if choice == 5 {
			break
		}
	}
}
