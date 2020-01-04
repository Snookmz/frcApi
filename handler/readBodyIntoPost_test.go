package handler

import (
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/databaseAbstraction"
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/logger"
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/loggerAbstraction"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReadBodyIntoPost(t *testing.T) {

	h := &Handler{}
	var err error

	l := logger.NewLogger(4, "../debug.log")
	h.Logger = loggerAbstraction.NewRegisteredLogger(l)

	d := &NsDatabase{}
	h.DbFunctions = databaseAbstraction.NewRegisteredDbFuncAbs(d)

	path := "/family/info"

	var p Post
	p.SessionId = "unit test session id"

	type Member struct {
		User string `json:"User"`
		Password string `json:"password"`
	}

	var m Member
	m.User = "admin"

	mBuf, err := json.Marshal(m)

	if err != nil {
		t.Errorf("error returned from json.Marshal: %s", err.Error())
		return
	}

	p.Json = mBuf

	var pBuf []byte
	pBuf, err = json.Marshal(p)
	if err != nil {
		t.Errorf("error returned from json.Marshal: %s", err.Error())
		return
	}

	var b *bytes.Buffer
	b = bytes.NewBuffer(pBuf)

	var r *http.Request
	r = httptest.NewRequest("POST", path, b)

	var rP Post
	rP, err = h.ReadBodyIntoPost(r)
	if err != nil {
		t.Errorf("error returned from ReadBody: %s\n", err.Error())
		return
	}

	var m2 Member
	err = json.Unmarshal(rP.Json, &m2)
	if err != nil {
		t.Errorf("error returned from json.Unmarshal: %s", err.Error())
		return
	}

	if rP.SessionId != p.SessionId {
		t.Errorf("expected session ids to match")
	}

	if m2.User != m.User || m2.Password != m.Password {
		t.Errorf("expected json to match")
	}








}
