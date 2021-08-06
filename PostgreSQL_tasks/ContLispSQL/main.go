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
type Contact struct {
	Id     int
	Name   string
	Gender string
	Phone  int
	Mail   string
}



func Menu() {
	fmt.Println("**************************")
	fmt.Println("*          Menu          *")
	fmt.Println("**************************")

	fmt.Println("Select a contact     - 1")
	fmt.Println("Add new contact      - 2")
	fmt.Println("Update a contact     - 3")
	fmt.Println("Delete a contact     - 4")
	fmt.Println("Exit                 - 5")
	fmt.Println("**************************")
}

func EnterDetails(c *Contact) {
	var phone int
	var name, gender, mail string

	fmt.Print("Enter name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter gender: ")
	fmt.Scanln(&gender)
	fmt.Print("Enter phone: ")
	fmt.Scanln(&phone)
	fmt.Print("Enter mail: ")
	fmt.Scanln(&mail)

	c.Name = name
	c.Gender = gender
	c.Phone = phone
	c.Mail = mail
}

func Add(c *Contact) {
	// Inserting
	sqlStatement := `
	INSERT INTO contacts (name, gender, phone, email)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, "John", "male", 5550004, "john@email.com").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}

func EnterDetails(c *Contact) {
	var phone int
	var name, gender, mail string

	fmt.Print("Enter name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter gender: ")
	fmt.Scanln(&gender)
	fmt.Print("Enter phone: ")
	fmt.Scanln(&phone)
	fmt.Print("Enter mail: ")
	fmt.Scanln(&mail)

	c.Name = name
	c.Gender = gender
	c.Phone = phone
	c.Mail = mail
}

func main() {

	// Creating the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

	// Opening and checking if we opened connector correctly
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var choice int
	var id int
	var cm ContactManager
	var c Contact

	for {
		Menu()
		fmt.Print("Enter a your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			cm.ContactList()
		} else if choice == 2 {
			c.Id = len(cm.contacts)
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
