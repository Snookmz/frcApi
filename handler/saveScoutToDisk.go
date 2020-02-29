package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func (h *Handler) SaveScoutsToDisk (scouts []Scout) (err error) {
	h.Logger.LPrintf(3, "SaveScoutsToDisk")

	for _, s := range scouts {
		h.Logger.LPrintf(3, "writing scout from '%s' device\n", s.ParentData.DeviceName)

		var scoutText []byte
		scoutText, err = json.Marshal(s)
		if err != nil {
			err = errors.New(fmt.Sprintf("error returned from json.Marshal: %s", err.Error()))
			return
		}

		var t time.Time
		t = time.Now()
		var tS string
		tS = t.Format(time.UnixDate)

		var fileName string
		fileName = fmt.Sprintf("%s %s %s  - %s.json", s.ParentData.DeviceName, s.ParentData.TeamDetails.IdTeam, s.ParentData.TxEvent, tS)

		//err = writeCsvFile(fileName, scoutText)

		err = h.SendFileToDropBox(fileName, scoutText)
		if err != nil {
			err = errors.New(fmt.Sprintf("error returned from SendFileToDropBox: %s", err.Error()))
			return
		}

	}

	return
}

func writeCsvFile(fileName string, data []byte) (err error) {

	//var w *csv.Writer
	//w = csv.NewWriter()

	return
}

func (h *Handler) SendFileToDropBox(fileName string, data []byte) (err error) {
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

	//{\"path\": \"/Homework/math/Matrices.txt\",\"mode\": \"add\",\"autorename\": true,\"mute\": false,\"strict_conflict\": false}"

	fmt.Printf("arg == %s\n", string(argByte))

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

/*
curl -X POST https://content.dropboxapi.com/2/files/upload \
    --header "Authorization: Bearer 9v6R9OUaq4cAAAAAAAAIEJl6nM3rcGIggQ2R-J7D-IQP1Vcjlsllgk83OyaiO84I" \
    --header "Dropbox-API-Arg: {\"path\": \"/Homework/math/Matrices.txt\",\"mode\": \"add\",\"autorename\": true,\"mute\": false,\"strict_conflict\": false}" \
    --header "Content-Type: application/octet-stream" \
    --data-binary @local_file.txt
 */
