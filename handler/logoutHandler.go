package handler

import (
	"errors"
	"fmt"
	"net/http"
)

func (h *Handler) LogoutHandler (w http.ResponseWriter, r *http.Request) {
	h.Logger.LPrintf(4, "LogoutHandler")

	err := h.DbFunctions.LogoutUser(h.User.Name)
	var j ReturnJson
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from LogoutUser: %s", err.Error()))
		h.Logger.LPrintf(1, err.Error())
		j.Success = false
		j.Message = err.Error()
		h.PrintOutput(w, j)
		return
	}

	j.Success = true
	j.Message = "user is not logged out"
	h.PrintOutput(w, j)
	return


}
