package nsc

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

func (n *Nsc) GetTopN (q Query) (d DataTopN, err error) {
	if q.Component != "stats" ||
		q.Action != "query" ||
		q.Item != "top_flows" ||
		q.Parameters.ClassId == 0 ||
		(q.Parameters.Direction != "inbound" && q.Parameters.Direction != "outbound") ||
		q.Parameters.Period == "" ||
		q.Parameters.Duration == 0 ||
		q.Parameters.Target == "" {

		var j string
		j, err = ReturnJsonFormat("topN")
		if err != nil {
			err = errors.New(fmt.Sprintf("error returned from returnJsonFormat: %s", err.Error()))
			return
		}
		err = errors.New(fmt.Sprintf("error topN query has incorrect format, expected: '%s', got %+v", j, q))
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

	d.Direction = q.Parameters.Direction

	err = json.Unmarshal(dB, &d)
	if err != nil {
		err = errors.New(fmt.Sprintf("error - no data, got '%s'", string(dB)))
		return
	}

	return

}
