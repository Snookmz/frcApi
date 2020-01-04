package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (d *NsDatabase) GetUserStateFromSessionId (sessionId string) (u UserState, err error) {

	var uSql UserStateSql
	err = d.Db.QueryRow("SELECT session_id, last_active, sign_on_time, signed_in, ip, name FROM ns_users_state WHERE session_id=$1", sessionId).Scan(&uSql.SessionId, &uSql.LastActive, &uSql.SignOnTime, &uSql.SignedIn, &uSql.Ip, &uSql.Name)

	if err == sql.ErrNoRows {
		err = errors.New("error no user with that session exists")
		return
	}

	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from db.QueryRow: %s", err.Error()))
		return
	}

	u = convertUserStateSqlToUserState(uSql)

	return

}
