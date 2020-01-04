package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelpHandler(t *testing.T) {

	var err error
	var p Post

	var post []byte
	post, err = json.Marshal(p)
	if err != nil {
		t.Errorf("error returned from json.Marshal: %s\n", err.Error())
	}

	req, err := http.NewRequest("POST", "/help", bytes.NewReader(post))
	if err != nil {
		t.Errorf("error returned from http.NewRequest: %s\n", err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	testServer().ServeHTTP(res, req)


	var result *http.Response
	result = res.Result()

	if result.StatusCode != 400 {
		t.Fatalf("expected status 400, got %v", result.StatusCode)
	}

}
