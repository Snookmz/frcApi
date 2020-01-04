package databaseAbstraction

import "bitbucket.org/turbosoftnetworks/netscope-api-v2/database"

type DbFunc interface {
	CreateClass (c database.Class) (newClass database.Class, err error)

	DeleteClass (classId int) (err error)

	EditClass (c database.Class) (err error)
	EditUser (userId int, u database.User) (err error)

	GetClasses (policyId int) (classes []database.Class, err error)

	LoginUser (name string, password string) (cookieId string, err error)
	LogoutUser (name string) (err error)
}

type RegisteredDbFuncAbs struct {
	DatabaseFunctions DbFunc
}

func NewRegisteredDbFuncAbs (dbFunc DbFunc) *RegisteredDbFuncAbs {
	return &RegisteredDbFuncAbs{dbFunc}
}

func (d RegisteredDbFuncAbs) CreateClass (c database.Class) (newClass database.Class, err error) {
	newClass, err = d.DatabaseFunctions.CreateClass(c)
	return
}

func (d RegisteredDbFuncAbs) EditClass (c database.Class) (err error) {
	err = d.DatabaseFunctions.EditClass(c)
	return
}



func (d RegisteredDbFuncAbs) LoginUser (name string, password string) (sessionId string, err error) {
	sessionId, err = d.DatabaseFunctions.LoginUser(name, password)
	return
}

func (d RegisteredDbFuncAbs) LogoutUser (name string) (err error) {
	err = d.DatabaseFunctions.LogoutUser(name)
	return
}
