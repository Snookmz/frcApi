package handler

import (
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/databaseAbstraction"
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/logger"
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/loggerAbstraction"
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/nscAbstraction"
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/nssAbstraction"
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/systemAbstraction"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}


func testServer () http.Handler {
	var h Handler

	l := logger.NewLogger(4, "../debug.log")
	h.Logger = loggerAbstraction.NewRegisteredLogger(l)

	d := &NsDatabase{}
	h.DbFunctions = databaseAbstraction.NewRegisteredDbFuncAbs(d)

	n :=&Nsc{}
	h.NscFunctions = nscAbstraction.NewRegisteredNscFuncAbs(n)

	ns := &Nss{}
	h.NssFunctions = nssAbstraction.NewRegisteredNssFuncAbs(ns)

	s := &NsSystem{}
	h.SystemFunctions = systemAbstraction.NewRegisteredSystemFuncAbs(s)

	var r *mux.Router
	testMw := func (w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		var p Post
		var err error
		p, err = h.ReadBodyIntoPost(r)
		if err != nil {
			fmt.Printf("error returned in testMw from ReadBodyIntoPostContext")
			return
		}
		ctx := context.WithValue(r.Context(), "PostContext", p)
		r = r.WithContext(ctx)
		next(w, r)
	}

	r = h.GenerateRoutes(testMw)

	return r
}

