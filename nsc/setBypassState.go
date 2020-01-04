package nsc

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

func (n *Nsc) SetBypassState (q Query) (d DataBypass, err error) {

	if q.Component != "bypass" ||
		q.Action != "command" ||
		(q.Command != "bypass" && q.Command != "exit_bypass") {
		var j string
		j, err = ReturnJsonFormat("bypass")
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
