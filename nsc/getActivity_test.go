package nsc

import (
	"github.com/fatih/pool"
	"testing"
)

func TestGetActivity (t *testing.T) {
	var d Data
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
	q.Item = "bandwidth_class"
	q.Parameters.ClassId = 8
	q.Parameters.Direction = "inbound"
	q.Parameters.Period = "current"
	q.Parameters.Duration = 300

	n := NewNsc(p)

	d, err = n.GetActivity(q)

	if err != nil {
		t.Errorf("error returned from GetActivity: %s", err.Error())
		return
	}

	if len(d.Result.Data) == 0 || len(d.Result.Time) == 0 {
		t.Errorf("expected data and time to be populated")
	}
}

func TestGetActivityDpi (t *testing.T) {
	var d Data
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
	q.Item = "bandwidth_class"
	q.Parameters.ClassId = 8
	q.Parameters.Direction = "inbound"
	q.Parameters.Period = "current"
	q.Parameters.Duration = 300
	q.Parameters.DpiProtocol = "Spotify"

	n := NewNsc(p)
	d, err = n.GetActivity(q)

	if err != nil {
		t.Errorf("error returned from GetActivity: %s", err.Error())
		return
	}

	if len(d.Result.Data) == 0 || len(d.Result.Time) == 0 {
		t.Errorf("expected data and time to be populated")
	}

}
