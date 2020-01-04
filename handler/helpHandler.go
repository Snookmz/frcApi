package handler

import "net/http"

func (h *Handler) HelpHandler(w http.ResponseWriter, r *http.Request) {
	h.Logger.LPrintf(4, "h.HelpHandler")
	var j ReturnJson
	j.Success = false
	j.Message = "not implemented"
	h.PrintOutput(w, j)

	h.Logger.LPrintf(4, "returning %+v", j)

	return
}
