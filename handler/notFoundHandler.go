package handler

import "net/http"

func (h *Handler) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	h.Logger.LPrintln(2, "NotFoundHandler")
	h.Logger.LPrintf(2,"Request Method: %+v", r.Method)
	h.Logger.LPrintf(2,"Requested URL : %+v", r.URL)
	h.Logger.LPrintf(2,"Authorization : %+v", r.Header["Authorization"])
	w.WriteHeader(http.StatusNotFound)

}
