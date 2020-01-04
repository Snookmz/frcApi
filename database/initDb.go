package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func InitDb(user string, password string, database string, host string) (db *sql.DB, err error) {
	// TODO timeout sql connections
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s",
		user, password, database, host)

	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		log.Panicf("error returned from sql.Open: %s\n", err.Error())
	}

	return

}
