package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (h *Handler) LoginHandler (w http.ResponseWriter, r *http.Request) {
	h.Logger.LPrintf(4, "h.LoginHandler")

	var j ReturnJson
	var err error

	var p Post
	p, err = h.ReadBodyIntoPost(r)
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from ReadPostFromRequest: %s", err.Error()))
		h.Logger.LPrintf(1, err.Error())
		j.Success = false
		j.Message = err.Error()
		h.PrintOutput(w, j)
		return
	}

	var l Login
	err = json.Unmarshal(p.Json, &l)
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from json.Unmarshal: %s", err.Error()))
		h.Logger.LPrintf(1, err.Error())
		j.Success = false
		j.Message = err.Error()
		h.PrintOutput(w, j)
		return
	}

	if l.Name == "" || l.Password == "" {
		err = errors.New(fmt.Sprintf("error - name and password required"))
		h.Logger.LPrintf(1, err.Error())
		j.Success = false
		j.Message = err.Error()
		h.PrintOutput(w, j)
		return
	}

	var sessionId string
	sessionId, err = h.DbFunctions.LoginUser(l.Name, l.Password)
	if err != nil {
		err = errors.New(fmt.Sprintf("user and password not accepted: %s", err.Error()))
		h.Logger.LPrintf(1, err.Error())
		j.Success = false
		j.Message = err.Error()
		h.PrintOutput(w, j)
		return
	}

	var s Session
	s.SessionId = sessionId
	j.Json, err = json.Marshal(s)
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from json.Unmarshal: %s", err.Error()))
		h.Logger.LPrintf(1, err.Error())
		j.Success = false
		j.Message = err.Error()
		h.PrintOutput(w, j)
		return
	}

	j.Success = true
	h.PrintOutput(w, j)
	return

}
