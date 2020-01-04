package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotFoundHandler(t *testing.T) {

	var err error
	var p Post

	var post []byte
	post, err = json.Marshal(p)
	if err != nil {
		t.Errorf("error returned from json.Marshal: %s\n", err.Error())
	}

	req, err := http.NewRequest("POST", "/notfound", bytes.NewReader(post))
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

	if result.StatusCode != 404 {
		t.Fatalf("expected status 404, got %v", result.StatusCode)
	}


}
