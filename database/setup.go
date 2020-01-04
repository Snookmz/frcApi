package database

import (
	"database/sql"
)

type NsDatabase struct {
	Db *sql.DB
}

func NewNsDatabase (db *sql.DB) *NsDatabase {
	return &NsDatabase{Db: db}
}

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	GroupId int `json:"group_id"`
	SiteId int `json:"site_id"`
	Active bool `json:"active"`
	StartDateMode string `json:"startDateMode"`
	StartDate string `json:"startDate"`
	StartTime string `json:"startTime"`
}

type UserSql struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	GroupId int `json:"group_id"`
	SiteId int `json:"site_id"`
	Active sql.NullBool `json:"active"`
	StartDateMode sql.NullString `json:"startDateMode"`
	StartDate sql.NullString `json:"startDate"`
	StartTime sql.NullString `json:"startTime"`
}

type UserState struct {
	SessionId string `json:"sessionId"`
	LastActive string `json:"lastActive"`
	SignOnTime string `json:"signOnTime"`
	SignedIn bool `json:"signedIn"`
	Ip string `json:"ip"`
	Name string `json:"name"`
}
type UserStateSql struct {
	SessionId string `json:"sessionId"`
	LastActive string `json:"lastActive"`
	SignOnTime sql.NullString `json:"signOnTime"`
	SignedIn sql.NullBool `json:"signedIn"`
	Ip sql.NullString `json:"ip"`
	Name string `json:"name"`
}
