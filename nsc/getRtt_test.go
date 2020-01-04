package nsc

import (
	"github.com/fatih/pool"
	"testing"
)

func TestGetRtt (t *testing.T) {
	var d DataRtt
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
	q.Item = "rtt"
	q.Parameters.ClassId = 8
	q.Parameters.Period = "current"
	q.Parameters.Duration = 300

	n := NewNsc(p)

	d, err = n.GetRtt(q)

	if err != nil {
		t.Errorf("error returned from GetRtt: %s", err.Error())
		return
	}

	if len(d.Result.Rtt) == 0 || len(d.Result.Time) == 0 || d.Status != "ok" {
		t.Errorf("expected ok, rtt and time to be populated")
	}

}
