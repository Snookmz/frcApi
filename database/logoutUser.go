package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (d *NsDatabase) LogoutUser(name string) (err error) {

	var stmt *sql.Stmt
	stmt, err = d.Db.Prepare("UPDATE ns_users_state SET signed_in=false WHERE name=$1 ")

	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from db.Prepare: %s", err.Error()))
		return
	}

	var res sql.Result
	res, err = stmt.Exec(name)

	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from stmt.Exec: %s", err.Error()))
		return
	}

	var rowsAffected int64
	rowsAffected, err = res.RowsAffected()
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from res.RowsAffected: %s", err.Error()))
		return
	}

	if rowsAffected != 1 {
		err = errors.New("less or more than 1 row affected")
		return
	}


	return
}
