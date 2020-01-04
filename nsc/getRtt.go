package nsc

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

func (n *Nsc) GetRtt(q Query) (d DataRtt, err error) {
	if q.Component != "stats" ||
		q.Action != "query" ||
		q.Item != "rtt" ||
		q.Parameters.ClassId == 0 ||
		q.Parameters.Duration == 0 {

		var j string
		j, err = ReturnJsonFormat("rtt")
		if err != nil {
			err = errors.New(fmt.Sprintf("error returned from returnJsonFormat: %s", err.Error()))
			return
		}
		err = errors.New(fmt.Sprintf("error rtt query has incorrect format, expected: %s", j))
		return
	}

	var qB []byte
	qB, err = json.Marshal(q)

	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from json.Marshal: %s", err.Error()))
		return
	}

	var dB []byte
	var conn net.Conn
	conn, err = n.Pool.Get()
	if err != nil {
		err = errors.New(fmt.Sprintf("error n.Pool.Get(): %s", err.Error()))
		return
	}
	dB, err = sendToConnection(conn, qB)
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from sendToConnection: %s", err.Error()))
		return
	}

	err = json.Unmarshal(dB, &d)
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from json.Unmarshal: %s", err.Error()))
		return
	}

	return

}
