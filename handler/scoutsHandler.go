package handler

import (
	"errors"
	"fmt"
	"net/http"
)

func (h *Handler) ScoutsHandler(w http.ResponseWriter, r *http.Request) {
	h.Logger.LPrintf(3, "ScoutsHandler")
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
}