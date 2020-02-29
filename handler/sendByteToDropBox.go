package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (h *Handler) SendByteToDropBox(fileName string, data []byte) (err error) {
	req, err := http.NewRequest("POST", h.DropBox.Url, bytes.NewReader(data))
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from http.NewRequest: %s", err.Error()))
		return
	}

	type Arg struct {
		Path string `json:"path"`
		Mode string `json:"mode"`
		AutoRename bool `json:"autorename"`
		Mute bool `json:"mute"`
		StrictConflict bool `json:"strict_conflict"`
	}

	var arg Arg
	arg.Path = "/5985/" + fileName
	arg.Mode = "add"
	arg.AutoRename = true
	arg.Mute = true
	arg.StrictConflict = false

	var argByte []byte
	argByte, err = json.Marshal(arg)
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from json.Marshal: %s", err.Error()))
		return
	}

	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Authorization", "Bearer " + h.DropBox.AccessToken)
	req.Header.Set("Dropbox-API-Arg", string(argByte))

	var client *http.Client
	client = &http.Client{}
	var resp *http.Response
	resp, err = client.Do(req)

	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from client.Do: %s", err.Error()))
		return
	}

	var content []byte
	content, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned form ioutil.ReadAll: %s", err.Error()))
		return
	}

	h.Logger.LPrintf(3, "content: %s", string(content))

	if strings.Contains(string(content), "Error") || strings.Contains(string(content), "error") {
		err = errors.New(fmt.Sprintf("error returned from POST: %s", string(content)))
		return
	}


	return
}
