package handler

import (
	"encoding/json"
	"errors"
	"fmt"
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

		err = h.SendByteToDropBox(fileName, scoutText)
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


/*
curl -X POST https://content.dropboxapi.com/2/files/upload \
    --header "Authorization: Bearer 9v6R9OUaq4cAAAAAAAAIEJl6nM3rcGIggQ2R-J7D-IQP1Vcjlsllgk83OyaiO84I" \
    --header "Dropbox-API-Arg: {\"path\": \"/Homework/math/Matrices.txt\",\"mode\": \"add\",\"autorename\": true,\"mute\": false,\"strict_conflict\": false}" \
    --header "Content-Type: application/octet-stream" \
    --data-binary @local_file.txt
 */
