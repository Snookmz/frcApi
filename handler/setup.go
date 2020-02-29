package handler

import (
	"encoding/json"
	"github.com/snookmz/frcApi/loggerAbstraction"
)

type Handler struct {
	Logger *loggerAbstraction.RegisteredLogger
	Dir string
	DropBox DropBox
}

type DropBox struct {
	AccessToken string `json:"accessToken"`
	Url string `json:"url"`
	Path string `json:"path"`
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
