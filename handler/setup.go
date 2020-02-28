package handler

import (
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/database"
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/databaseAbstraction"
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/loggerAbstraction"
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/nscAbstraction"
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/nssAbstraction"
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/systemAbstraction"
	"encoding/json"
)

type Handler struct {
	Logger *loggerAbstraction.RegisteredLogger
	DbFunctions *databaseAbstraction.RegisteredDbFuncAbs
	NscFunctions *nscAbstraction.RegisteredNscFuncAbs
	NssFunctions *nssAbstraction.RegisteredNssFuncAbs
	SystemFunctions *systemAbstraction.RegisteredSystemFuncAbs
	User database.User
	UserState database.UserState
}

type ReturnJson struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Json json.RawMessage `json:"json"`
}


type Email struct {
	Support string
	Sales string
	Info string
}

type Post struct {
	SessionId string `json:"sessionId"`
	Json json.RawMessage `json:"json"`
}

type Login struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

type Session struct {
	SessionId string `json:"id"`
}

type Password struct {
	Name string `json:"name"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
