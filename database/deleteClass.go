package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (d *NsDatabase) DeleteClass (classId int) (err error) {

	var stmt *sql.Stmt
	stmt, err = d.Db.Prepare("DELETE FROM ns_policy_class WHERE class_id=$1")
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from db.Prepare: %s", err.Error()))
		return
	}

	defer stmt.Close()
	var res sql.Result
	res, err = stmt.Exec(classId)
		if err != nil {
		err = errors.New(fmt.Sprintf("error returned from stmt.Exec: %s", err.Error()))
		return
	}

	var rowCnt int64
	rowCnt, err = res.RowsAffected()
	if rowCnt != 1 {
		err = errors.New(fmt.Sprintf("more or less than one row affected, rows: %d", rowCnt))
		return
	}

	return
}
