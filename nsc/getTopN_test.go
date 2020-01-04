package nsc

import (
	"github.com/fatih/pool"
	"testing"
)

func TestGetTopN (t *testing.T) {
	var d DataTopN
	var err error
	var q Query

	var p pool.Pool
	p, err = CreateConnectionPool("10.0.1.14", 4213)
	if err != nil {
		t.Errorf("error CreateConnectionPool: %s", err.Error())
		return
	}

	q.Component = "stats"
	q.Action = "query"
	q.Item = "top_flows"
	q.Parameters.ClassId = 8
	q.Parameters.Direction = "inbound"
	q.Parameters.Period = "current"
	q.Parameters.Duration = 300

	n := NewNsc(p)

	d, err = n.GetTopN(q)

	if err == nil {
		t.Errorf("expected error")
	}

	q.Parameters.Target = "ip_src"
	d, err = n.GetTopN(q)

	if err != nil {
		t.Errorf("error returned from GetTopN: %s", err.Error())
		return
	}

	if len(d.Result.Targets) == 0 || len(d.Result.Volumes) == 0 || d.Status != "ok" {
		t.Errorf("expected data and time to be populated")
	}

}
