package handler

import (
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/database"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (h *Handler) CheckSessionHandler (w http.ResponseWriter, r *http.Request) {
	h.Logger.LPrintf(4, "h.CheckSessionHandler")

	var j ReturnJson
	var err error

	var p Post
	p, err = ReadPostFromRequest(r)
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from ReadPostFromRequest: %s", err.Error()))
		h.Logger.LPrintf(1, err.Error())
		j.Success = false
		j.Message = err.Error()
		h.PrintOutput(w, j)
		return
	}

	var u database.UserState
	u, err = h.DbFunctions.GetUserStateFromSessionId(p.SessionId)

	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from GetUserStateFromSessionId: %s", err.Error()))
		h.Logger.LPrintf(1, err.Error())
		j.Success = false
		j.Message = err.Error()
		h.PrintOutput(w, j)
		return
	}

	j.Json, err = json.Marshal(u)
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from json.Marshal: %s", err.Error()))
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
