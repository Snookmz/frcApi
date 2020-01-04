package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

func (d *NsDatabase) CreateClass (c Class) (newClass Class, err error) {
	var nsId string
	nsId = strconv.Itoa(c.NsId)

	var stmt *sql.Stmt
	stmt, err = d.Db.Prepare(`INSERT INTO ns_policy_class 
    (ns_id, 
     name, 
     parent_id, 
     prev_id, 
     next_id, 
     first_child_id, 
     last_child_id, 
     precedence, 
     policy_id
     ) 
     VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
     RETURNING 
     ns_id, 
      class_id, 
      name, 
      parent_id, 
      prev_id, 
      next_id, 
      first_child_id, 
      last_child_id, 
      precedence, 
      policy_id`)

	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from db.Prepare: %s", err.Error()))
		return
	}

	var newNsId string

	err = stmt.QueryRow(nsId, c.Name, c.ParentId, c.PrevId, c.NextId, c.FirstChildId, c.LastChildId, c.Precedence, c.PolicyId).Scan(&newNsId, &newClass.ClassId, &newClass.Name, &newClass.ParentId, &newClass.PrevId, &newClass.NextId, &newClass.FirstChildId, &newClass.LastChildId, &newClass.Precedence, &newClass.PolicyId)
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from stmt.QueryRow: %s", err.Error()))
		return
	}

	newClass.NsId, err = strconv.Atoi(newNsId)
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from strconv.Atoi: %s", err.Error()))
		return
	}

	return
}

