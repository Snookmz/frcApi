package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
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

		// [EVENT]_[MATCH or PIT]_[TEAM_No]_[DEVICE]_[YYYYMMDD_HHMMSS]

		var event = strings.ToUpper(p.Event)
		var device = strings.ToUpper(p.Record.TxComputerName)
		if device == "" {
			device = "UNKNOWN"
		}

		var fileName string
		fileName = fmt.Sprintf("%s_PIT_%v_%s_%s.json", event, p.Details.IdTeam, device, p.Record.DtCreated)

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

