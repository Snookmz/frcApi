package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (d *NsDatabase) GetUsers () (users []User, err error) {

	var rows *sql.Rows
	rows, err = d.Db.Query("SELECT id, name, group_id, site_id, active, startdatemode, startdate, starttime FROM ns_users")
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from db.Query: %s", err.Error()))
		return
	}

	for rows.Next() {
		var uSql UserSql
		err = rows.Scan(&uSql.Id, &uSql.Name, &uSql.GroupId, &uSql.SiteId, &uSql.Active, &uSql.StartDateMode, &uSql.StartDate, &uSql.StartTime)
		if err != nil {
			err = errors.New(fmt.Sprintf("error returned from rows.Scan: %s", err.Error()))
			return
		}
		users = append(users, convertUserSqlToUser(uSql))
	}


	return
}
