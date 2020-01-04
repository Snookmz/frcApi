package nsc

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

func (n *Nsc) GetActivity(q Query) (d Data, err error) {

	if q.Component != "stats" ||
		q.Action != "query" ||
		q.Item != "bandwidth_class" ||
		q.Parameters.ClassId == 0 ||
		(q.Parameters.Direction != "inbound" && q.Parameters.Direction != "outbound") ||
		q.Parameters.Period == "" ||
		q.Parameters.Duration == 0 {

		var j string
		j, err = ReturnJsonFormat("activity")
		if err != nil {
			err = errors.New(fmt.Sprintf("error returned from returnJsonFormat: %s", err.Error()))
			return
		}
		err = errors.New(fmt.Sprintf("error activity query has incorrect format, expected: '%s' got '%+v", j, q))
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

	if string(dB) == "" {
		err = errors.New(fmt.Sprintf("error, nothing returned from nsc, sent - %s", string(qB)))
	}

	err = json.Unmarshal(dB, &d)
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from json.Unmarshal: %s", err.Error()))
		return
	}

	d.Direction = q.Parameters.Direction

	if d.Result.Error != "" {
		err= errors.New(fmt.Sprintf("error returned from nsc '%s', sent '%s'", d.Result.Error, string(qB)))
	}

	return
}
