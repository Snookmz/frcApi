package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

func (h *Handler) NsAuthenticationMw (w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	h.Logger.LPrintf(4, "NsAuthenticationMw URL: %s", r.URL)

	var j ReturnJson
	var err error
	var p Post

	p, err = h.ReadBodyIntoPost(r)

	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from h.ReadBodyIntoPost: %s", err.Error()))
		h.Logger.LPrintf(1, err.Error())
		j.Success = false
		j.Message = err.Error()
		h.PrintOutput(w, j)
		return
	}

	ctx := context.WithValue(r.Context(), "PostContext", p)
	r = r.WithContext(ctx)

	if p.SessionId == "" {
		err = errors.New("error - no sessionId sent")
		h.Logger.LPrintf(1, err.Error())
		j.Success = false
		j.Message = err.Error()
		h.PrintOutput(w, j)
		return
	}

	h.Logger.LPrintf(4, "post: %+v", p)

	h.UserState, err = h.DbFunctions.GetUserStateFromSessionId(p.SessionId)

	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from GetUserStateFromSessionId: %s", err.Error()))
		h.Logger.LPrintf(1, err.Error())
		j.Success = false
		j.Message = err.Error()
		h.PrintOutput(w, j)
		return
	}

	if !h.UserState.SignedIn {
		h.Logger.LPrintf(2, "user not logged in")
		j.Success = false
		j.Message = "user not logged in"
		h.PrintOutput(w, j)
		return
	}

	h.Logger.LPrintf(4, "complete, passing on to handler for URL: %s", r.URL)
	next(w, r)


}
