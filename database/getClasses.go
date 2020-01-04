package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (d *NsDatabase) GetClasses(policyId int) (classes []Class, err error) {
	var rows *sql.Rows
	rows, err = d.Db.Query(`select 
       ns_id, 
       class_id, 
       name, 
       parent_id, 
       prev_id, 
       next_id, 
       first_child_id,
       last_child_id,
       precedence,
       policy_id 
	from ns_policy_class 
	WHERE 
	      policy_id=$1`, policyId)

	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from db.Query: %s", err.Error()))
		return
	}

	for rows.Next() {
		var c Class
		err = rows.Scan(&c.NsId, &c.ClassId, &c.Name, &c.ParentId, &c.PrevId, &c.NextId, &c.FirstChildId, &c.LastChildId, &c.Precedence, &c.PolicyId)
		if err != nil {
			err = errors.New(fmt.Sprintf("error returned from rows.Scan: %s", err.Error()))
			return
		}
		classes = append(classes, c)
	}
	return
}

