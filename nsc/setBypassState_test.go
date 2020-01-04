package nsc

import (
	"github.com/fatih/pool"
	"testing"
)

func TestSetBypassStateOnline (t *testing.T) {
	var d DataBypass
	var err error
	var q Query

	var p pool.Pool
	p, err = CreateConnectionPool("10.0.1.14", 4213)
	if err != nil {
		t.Errorf("error CreateConnectionPool: %s", err.Error())
		return
	}

	q.Component = "bypass"
	q.Action = "command"
	q.Command = "exit_bypass"

	n := NewNsc(p)

	d, err = n.SetBypassState(q)

	if err != nil {
		t.Errorf("error returned from GetActivity: %s", err.Error())
		return
	}

	if d.Result.State != "Online" {
		t.Errorf("expected state to be bypass")
	}

}

func TestSetBypassStateOffline (t *testing.T) {
	var d DataBypass
	var err error
	var q Query
	var p pool.Pool
	p, err = CreateConnectionPool("10.0.1.14", 4213)
	if err != nil {
		t.Errorf("error CreateConnectionPool: %s", err.Error())
		return
	}

	q.Component = "bypass"
	q.Action = "command"
	q.Command = "bypass"

	n := NewNsc(p)

	d, err = n.SetBypassState(q)

	if err != nil {
		t.Errorf("error returned from GetActivity: %s", err.Error())
		return
	}

	if d.Result.State != "Bypass active" {
		t.Errorf("expected state to be bypass")
	}

}
