package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginHandler (t *testing.T) {
	var err error
	var p Post

	var l Login
	l.Name = "unit test name"
	l.Password = "unit test password"

	p.Json, err = json.Marshal(l)
	if err != nil {
		t.Errorf("error returned from json.Marshal: %s", err.Error())
		return
	}

	var post []byte
	post, err = json.Marshal(p)
	if err != nil {
		t.Errorf("error returned from json.Marshal: %s\n", err.Error())
	}

	req, err := http.NewRequest("POST", "/login", bytes.NewReader(post))
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

	var s Session
	err = json.Unmarshal(r.Json, &s)
		if err != nil {
		t.Errorf("error returned from json.Unmarshal: %s", err.Error())
		return
	}

	if s.SessionId == "" {
		t.Errorf("expected session id to be populated")
	}
}
func TestLoginHandlerFail (t *testing.T) {
	var err error
	var p Post

	var l Login

	p.Json, err = json.Marshal(l)
	if err != nil {
		t.Errorf("error returned from json.Marshal: %s", err.Error())
		return
	}

	var post []byte
	post, err = json.Marshal(p)
	if err != nil {
		t.Errorf("error returned from json.Marshal: %s\n", err.Error())
	}

	req, err := http.NewRequest("POST", "/login", bytes.NewReader(post))
	if err != nil {
		t.Errorf("error returned from http.NewRequest: %s\n", err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "bearer blah")
	res := httptest.NewRecorder()
	testServer().ServeHTTP(res, req)


	var result *http.Response
	result = res.Result()

	if result.StatusCode != 400 {
		t.Fatalf("expected status 400, got %v", result.StatusCode)
	}

	var r ReturnJson
	err = json.Unmarshal(res.Body.Bytes(), &r)
	if err != nil {
		t.Errorf("error returned from json.Unmarshal: %s", err.Error())
		return
	}

	if r.Success {
		t.Errorf("expected failure")
	}
}
