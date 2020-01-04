package nsc

import (
	"encoding/json"
	"net"
	"testing"
)

func TestCreateConnection (t *testing.T) {
	var err error
	var conn net.Conn


	conn, err = CreateConnection("10.0.1.14", 4213)
	if err != nil {
		t.Errorf("error returned from createConnection: %s", err.Error())
		return
	}

	defer conn.Close()

	var q Query
	q.Component = "stats"
	q.Action = "query"
	q.Item = "bandwidth_class"
	q.Parameters.ClassId = 8
	q.Parameters.Direction = "inbound"
	q.Parameters.Period = "current"
	q.Parameters.Duration = 300

	var data []byte
	data, err = json.Marshal(q)

	if err != nil {
		t.Errorf("error returned from json.Marshal: %s", err.Error())
	}

	var returnData []byte
	returnData, err = sendToConnection(conn, data)

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
