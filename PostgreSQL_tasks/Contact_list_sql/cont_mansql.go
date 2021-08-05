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
	INSERT INTO contacts (name, gender, phone, email)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, "Brian", "male", 5550003, "briangr@gmail.com").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)

	/*
		// Actually opening up a connection
		err = db.Ping()
		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully connected.")
	*/
}

