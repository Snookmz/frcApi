package nsc

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

func (n *Nsc) GetBypassState () (d DataBypass, err error) {
	var q Query
	q.Component = "bypass"
	q.Action = "query"
	q.Item = "bypass"

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

	if string(dB) == "" {
		err = errors.New(fmt.Sprintf("error, nothing returned from nsc, sent - %s", string(qB)))
	}

	err = json.Unmarshal(dB, &d)
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from json.Unmarshal: %s", err.Error()))
		return
	}
	return
}
