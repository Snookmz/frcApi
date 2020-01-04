package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func (h *Handler) ReadBodyIntoPost(r *http.Request) (p Post, err error) {
	h.Logger.LPrintln(4,"ReadBodyIntoPostContext")

	var body []byte
	if r.Body == nil {
		err = errors.New("error: body is nil")
		h.Logger.LPrintf(1, err.Error())
		return
	}
	body, err = ioutil.ReadAll(r.Body)

	if err != nil {
		h.Logger.LPrintf(1, "error returned from ioutil.ReadAll: %s\n", err)
		return
	}

	defer r.Body.Close()

    h.Logger.LPrintf(3, "REQUEST: %s, body: %s", r.URL.Path, string(body))

	err = json.Unmarshal(body, &p)
	if err != nil {
		h.Logger.LPrintf(1, "error returned from json.Unmarshal: %s\n", err)
		if e, ok := err.(*json.SyntaxError); ok {
			h.Logger.LPrintf(1,"syntax error at byte offset %d", e.Offset)
		}
		h.Logger.LPrintf(1, "body: %q", body)
		return
	}

	return

}
