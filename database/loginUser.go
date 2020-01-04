package database

import (
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func (d *NsDatabase) LoginUser (name string, password string) (sessionId string, err error) {
	var dbPassword string
	err = d.Db.QueryRow("SELECT password FROM ns_users WHERE name=$1", name).Scan(&dbPassword)

	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from QueryRow: %s", err.Error()))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
		if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword || err == bcrypt.ErrHashTooShort {
			err = errors.New("password incorrect")
		} else {
			err = errors.New(fmt.Sprintf("error returned from bcrypt.CompareHashAndPassword: %s", err.Error()))
		}
		return
	}

	// ON CONFLICT doesn't exist in this version of Postgres, have to split the queries up
	var dbName string
	err = d.Db.QueryRow("SELECT name FROM ns_users_state WHERE name=$1", name).Scan(&dbName)
	if err != nil && err != sql.ErrNoRows {
		err = errors.New(fmt.Sprintf("error returned from QueryRow: %s", err.Error()))
		return
	}

	if dbName != "" {
		err = d.Db.QueryRow("UPDATE ns_users_state SET (sign_on_time, signed_in) = (current_timestamp, 'true') WHERE name=$1 RETURNING session_id", name).Scan(&sessionId)
	} else {
		err = d.Db.QueryRow("INSERT INTO ns_users_state (name, sign_on_time, signed_in) VALUES ($1, current_timestamp, 'true') RETURNING session_id", name).Scan(&sessionId)
	}

	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from db.QueryRow: %s", err.Error()))
		return
	}

	return
}
