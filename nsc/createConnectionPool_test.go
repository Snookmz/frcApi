package nsc

import (
	"encoding/json"
	"github.com/fatih/pool"
	"net"
	"testing"
)

func TestCreateConnectionPool(t *testing.T) {
	var err error
	var p pool.Pool

	p, err = CreateConnectionPool("10.0.1.14", 4213)
	if err != nil {
		t.Errorf("error returned from CreateConnectionPool: %s", err.Error())
		return
	}

	var q Query
	q.Component = "stats"
	q.Action = "query"
	q.Item = "bandwidth_class"
	q.Parameters.ClassId = 8
	q.Parameters.Direction = "inbound"
	q.Parameters.Period = "current"
	q.Parameters.Duration = 300
	q.Parameters.IpSrc = "0.0.0.0"
	q.Parameters.IpDst = "0.0.0.0"

	var data []byte
	data, err = json.Marshal(q)

	if err != nil {
		t.Errorf("error returned from json.Marshal: %s", err.Error())
	}

	for i := 0; i < 100; i++ {
		var returnData []byte
		var conn net.Conn
		conn, err = p.Get()
		if err != nil {
			t.Errorf("error returned from p.Get: %s", err.Error())
			return
		}

		returnData, err = sendToConnection(conn, data)

		if err != nil {
			t.Errorf("error returned from sendToConnection: %s", err.Error())
			return
		}

		//fmt.Printf("return: %s\n", string(returnData))

		var result Data
		err = json.Unmarshal(returnData, &result)

		if err != nil {
			t.Errorf("error returned from json.Unmarshal: %s", err.Error())
		}

		if len(result.Result.Data) == 0 {
			t.Errorf("expected data to be populated")
		}

		if len(result.Result.Time) == 0 {
			t.Errorf("expected time to be populated")
		}
	}

}
