package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func (h *Handler) PitsHandler(w http.ResponseWriter, r *http.Request) {
	h.Logger.LPrintf(3, "PitsHandler")

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

	var pits []Pit
	err = json.Unmarshal(p.Json, &pits)
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from json.Unmarshal: %s", err.Error()))
		h.Logger.LPrintf(1, err.Error())
		j.Success = false
		j.Message = err.Error()
		h.PrintOutput(w, j)
		return
	}

	h.Logger.LPrintf(4, "Pits: %+v\n", pits)

	for _, p := range pits {
		var t time.Time
		t = time.Now()
		var tS string
		tS = t.Format(time.UnixDate)

		var fileName string
		fileName = fmt.Sprintf("Pit for %s (%s).json", p.Details.IdTeam, tS)

		b, err := json.Marshal(p)
		if err != nil {
			err = errors.New(fmt.Sprintf("error returned from json.Marshal: %s", err.Error()))
			h.Logger.LPrintf(1, err.Error())
			j.Success = false
			j.Message = err.Error()
			h.PrintOutput(w, j)
			return
		}

		err = h.SendByteToDropBox(fileName, b)
		if err != nil {
			err = errors.New(fmt.Sprintf("error returned from SendByteToDropBox: %s", err.Error()))
			h.Logger.LPrintf(1, err.Error())
			j.Success = false
			j.Message = err.Error()
			h.PrintOutput(w, j)
			return
		}
	}

	j.Success = true
	h.PrintOutput(w, j)
	return

}
