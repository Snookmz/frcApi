package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (h *Handler) ScoutsHandler(w http.ResponseWriter, r *http.Request) {
	h.Logger.LPrintf(3, "ScoutsHandler")
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

	var scouts []Scout
	err = json.Unmarshal(p.Json, &scouts)
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from json.Unmarshal: %s", err.Error()))
		h.Logger.LPrintf(1, err.Error())
		j.Success = false
		j.Message = err.Error()
		h.PrintOutput(w, j)
		return
	}

	h.Logger.LPrintf(4, "Scouts: %+v\n", scouts)

	//var scoutsCsv []ScoutCsv
	//scoutsCsv = convertScoutsToScoutsCsv(scouts)
	//
	//var csv string
	//csv, err = gocsv.MarshalString(&scoutsCsv)
	//if err != nil {
	//	err = errors.New(fmt.Sprintf("error returned from csv.MarshalString: %s", err.Error()))
	//	h.Logger.LPrintf(1, err.Error())
	//	j.Success = false
	//	j.Message = err.Error()
	//	h.PrintOutput(w, j)
	//	return
	//}
	//
	//h.Logger.LPrintf(3, "csv: %s", csv)

	err = h.SaveScoutsToDisk(scouts)
	if err != nil {
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