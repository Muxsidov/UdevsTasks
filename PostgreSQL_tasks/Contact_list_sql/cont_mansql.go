// Code is from https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/

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

func main() {

	// Create struct for SELECT *
	type Contact struct {
		ID int
		Name string 
		Gender string 
		Number int 
		Email string
		
	}	
	
	// Creating the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Opening and checking if we opened connector correctly
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	sqlStatement := `SELECT * FROM contacts WHERE id=$1;`
	var contact Contact 
	row := db.QueryRow(sqlStatement, 4)
	err_0 := row.Scan(&contact.ID, &contact.Name, &contact.Gender,
		&contact.Number, &contact.Email)
	switch err_0 {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
			return	
		case nil:
			fmt.Println(contact)
		default:
			panic(err_0)	
	}

	/*
	// Updating 
	sqlStatement := `
	UPDATE contacts
	SET name = $2, gender = $3, phone = $4, email = $5
	WHERE id = $1
	RETURNING id, name;`
	//res, err := db.Exec(sqlStatement, 3, "Changed name", "chander", 1234, "changedmail@mail.com")
	var name string
	var id int
	err = db.QueryRow(sqlStatement, 2, "Changed name", "chander", 1234, "changedmail@mail.com").Scan(&id, &name)
	if err != nil {
		panic(err)
	}
	fmt.Println(id, name)	
	*/
	
	/*
	// Deleting 
	sqlStatement := `
	DELETE FROM contacts
	WHERE id = $1;`
	res, err := db.Exec(sqlStatement, 3)
	if err != nil {
		panic(err)
	}
	
	// Cheking if Updating/Deleting was succsessful 
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
	*/
	
	/*
	// Inserting
	sqlStatement := `
	INSERT INTO contacts (name, gender, phone, email)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, "Brian", "male", 5550003, "briangr@gmail.com").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
	*/

	/*
		// Actually opening up a connection
		err = db.Ping()
		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully connected.")
	*/
}

