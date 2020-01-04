package nsc

import (
	"github.com/fatih/pool"
	"testing"
)

func TestGetBypassState (t *testing.T) {
	var d DataBypass
	var err error

	var p pool.Pool
	p, err = CreateConnectionPool("10.0.1.14", 4213)
	if err != nil {
		t.Errorf("error CreateConnectionPool: %s", err.Error())
		return
	}

	n := NewNsc(p)
	d, err = n.GetBypassState()

	if err != nil {
		t.Errorf("error returned from GetBypassState: %s", err.Error())
		return
	}

	if d.Status != "ok" || d.Result.State == "" {
		t.Errorf("expected status ok and state not to == ''")
	}

}
