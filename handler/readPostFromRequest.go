package handler

import (
	"errors"
	"net/http"
)

func ReadPostFromRequest(r *http.Request) (p Post, err error) {
	var ok bool
	pI := r.Context().Value("PostContext")
	p, ok = pI.(Post)

	if !ok {
		err = errors.New("error converting post interface to Post")
		return
	}
	return
}