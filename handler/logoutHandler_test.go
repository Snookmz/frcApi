package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogoutHandler (t *testing.T) {
	var err error
	var p Post

	var post []byte
	post, err = json.Marshal(p)
	if err != nil {
		t.Errorf("error returned from json.Marshal: %s\n", err.Error())
	}

	req, err := http.NewRequest("POST", "/user/logout", bytes.NewReader(post))
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
