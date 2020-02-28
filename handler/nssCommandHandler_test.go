package handler

import (
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/nss"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSystemNetworkHandler (t *testing.T) {
	var err error
	var p Post

	var q nss.Query
	p.Json, err = json.Marshal(q)

	if err != nil {
		t.Errorf("error returned from json.Marshal")
	}
	var post []byte
	post, err = json.Marshal(p)
	if err != nil {
		t.Errorf("error returned from json.Marshal: %s\n", err.Error())
	}

	req, err := http.NewRequest("POST", "/nss", bytes.NewReader(post))
	if err != nil {
		t.Errorf("error returned from http.NewRequest: %s\n", err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()
	testServer().ServeHTTP(res, req)

	var result *http.Response
	result = res.Result()

	if result.StatusCode != 200 {
		t.Fatalf("expected status 200, got %v", result.StatusCode)
	}

	var r ReturnJson
	err = json.Unmarshal(res.Body.Bytes(), &r)
	if err != nil {
		t.Errorf("error returned from json.Unmarshal: %s", err.Error())
		return
	}

	if !r.Success {
		t.Errorf("expected success")
	}

}