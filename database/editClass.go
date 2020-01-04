package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

func (d *NsDatabase) EditClass (c Class) (err error) {
	var nsId string
	nsId = strconv.Itoa(c.NsId)

	var stmt *sql.Stmt
	stmt, err = d.Db.Prepare(`
	UPDATE ns_policy_class 
	SET (
		ns_id, 
		name, 
		parent_id, 
		prev_id, 
		next_id, 
		first_child_id, 
		last_child_id, 
		precedence,
	    policy_id
    ) = ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	WHERE class_id=$10`)

	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from db.Prepare: %s", err.Error()))
		return
	}

	defer stmt.Close()
	var res sql.Result
	res, err = stmt.Exec(nsId, c.Name, c.ParentId, c.PrevId, c.NextId, c.FirstChildId, c.LastChildId, c.Precedence, c.PolicyId, c.ClassId)
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
