package handler

import (
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/nss"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (h *Handler) NssCommandHandler (w http.ResponseWriter, r *http.Request) {
	h.Logger.LPrintf(3, "NssCommandHandler")

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

	var q nss.Query
	err = json.Unmarshal(p.Json, &q)
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from json.Unmarshal: %s", err.Error()))
		h.Logger.LPrintf(1, err.Error())
		j.Success = false
		j.Message = err.Error()
		h.PrintOutput(w, j)
		return
	}

	h.Logger.LPrintf(4, "Query: %+v\n", q)

	var nssReturn nss.NssReturn
	nssReturn, err = h.NssFunctions.SendCommandToNss(q)
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from SendCommandToNss: %s", err.Error()))
		h.Logger.LPrintf(1, err.Error())
		j.Success = false
		j.Message = err.Error()
		h.PrintOutput(w, j)
		return
	}

	j.Json, err = json.Marshal(nssReturn)
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
