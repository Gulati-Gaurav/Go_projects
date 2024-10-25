package main

import (
	"database/sql"
	"fmt"
	"log"
	// This is how to include package and also not use it
	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=gaurav dbname=postgres password=data@23 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM phnumbers")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var phNumbers []string
		if err := rows.Scan(&phNumbers); err != nil {
			log.Fatal(err)
		}
		fmt.Println(phNumbers)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
