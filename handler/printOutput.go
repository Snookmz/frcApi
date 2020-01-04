package handler

import (
	"encoding/json"
	"io"
	"net/http"
)

func (h *Handler) PrintOutput(w http.ResponseWriter, returnJ ReturnJson) {

	var err error
	var j []byte

	j, err = json.Marshal(returnJ)
	if err != nil {
		var	errJson = "{\"success\":false,\"message\":\"error creating return Json\"}"
		h.Logger.LPrintf(1, "error returned from json.Marshal: %s", err.Error())
		h.Logger.LPrintln(1, "RESPONSE:", errJson)
		_, _ = io.WriteString(w, errJson)

	} else {
		h.Logger.LPrintf(3, "RESPONSE: %s", string(j))
		if !returnJ.Success {
			w.WriteHeader(400)
		}
		_, _ = io.WriteString(w, string(j))
	}

	return
}
