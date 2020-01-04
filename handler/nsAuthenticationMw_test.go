package handler

import (
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/databaseAbstraction"
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/logger"
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/loggerAbstraction"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNsAuthenticationMw (t *testing.T) {

	var err error
	var p Post
	p.SessionId = "unit test session id"

	var post []byte
	post, err = json.Marshal(p)
	if err != nil {
		t.Errorf("error returned from json.Marshal: %s", err.Error())
		return
	}

	var h Handler

	l := logger.NewLogger(4, "../debug.log")
	h.Logger = loggerAbstraction.NewRegisteredLogger(l)

	d := &NsDatabase{}
	h.DbFunctions = databaseAbstraction.NewRegisteredDbFuncAbs(d)

	r := httptest.NewRequest("POST", "/policy/list", bytes.NewReader(post))

	w := httptest.NewRecorder()

	h.NsAuthenticationMw(w, r, h.testHandler)

	var resp *http.Response
	resp = w.Result()

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)

	var returnJ ReturnJson
	err = json.Unmarshal(body, &returnJ)

	if !returnJ.Success {
		t.Errorf("expected success, message: %s", returnJ.Message)
	}


}

func (h *Handler) testHandler (w http.ResponseWriter, r *http.Request) {
	var j ReturnJson
	j.Success = true
	j.Message = "testHandler"
	h.PrintOutput(w, j)
	return
}
