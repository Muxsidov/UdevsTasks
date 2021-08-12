// https://www.calhoun.io/using-postgresql-with-go/

package main

import (
	"database/sql"
	"fmt"

	//"github.com/jmoiron/sqlx"
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

type ContactManager struct {
	db *sql.DB
}

// Create a struct for SELECT *
type Contact struct {
	Id     int
	Name   string
	Gender string
	Phone  int
	Mail   string
}

func NewContactManager() (ContactManager, error) {
	cm := ContactManager{}
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

	fmt.Println("List of contacts     - 1")
	fmt.Println("Add new contact      - 2")
	fmt.Println("Update a contact     - 3")
	fmt.Println("Delete a contact     - 4")
	fmt.Println("Exit                 - 5")
	fmt.Println("**************************")
}

func (cm *ContactManager) Add(c *Contact) error {
	// func Add(c *Contact) error{

	// Inserting
	sqlStatement := `
	INSERT INTO contacts (name, gender, phone, email)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 0
	cm.db.QueryRow(sqlStatement, &c.Name, &c.Gender, &c.Phone, &c.Mail).Scan(&id)
	//	err = db.QueryRow(sqlStatement, &c.Name, &c.Gender, &c.Phone, &c.Mail).Scan(&id)
	//	if err != nil {
	//		panic(err)
	//	}
	fmt.Println("New record ID is:", id)
	return nil
}

func (cm *ContactManager) Update(c *Contact) {
	var err error

	// Updating
	sqlStatement := `
	UPDATE contacts
	SET name = $2, gender = $3, phone = $4, email = $5
	WHERE id = $1
	RETURNING id, name;`
	var name string
	var id int
	err = cm.db.QueryRow(sqlStatement, &c.Id, &c.Name, &c.Gender, &c.Phone, &c.Mail).Scan(&id, &name)
	if err != nil {
		panic(err)
	}
	fmt.Println(id, name)
}

func (cm *ContactManager) Delete(c *Contact) {
	// Deleting
	sqlStatement := `
	DELETE FROM contacts
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

func (cm *ContactManager) ContactList() {

	rows, err := cm.db.Query("SELECT id, name, phone, email FROM contacts LIMIT $1", 3)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, phone int
		var name, email string
		err = rows.Scan(&id, &name, &phone, &email)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(id, name, phone, email)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
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
