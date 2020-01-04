package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func PrepareTestDatabase() (db *sql.DB) {

	var err error
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s",
		"postgres", "n5admin", "netscope_test", "10.0.1.14")

	db, err = sql.Open("postgres", dbInfo )
	if err != nil {
		log.Printf("error returned from sql.Open: %s", err.Error())
		log.Fatal(err)
	}

	return
}
