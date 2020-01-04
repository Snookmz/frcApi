package nsc

import (
	"github.com/fatih/pool"
	"testing"
)

func TestGetLicense (t *testing.T) {
	var d DataLicense
	var err error
	var q Query

	var p pool.Pool
	p, err = CreateConnectionPool("10.0.1.14", 4213)
	if err != nil {
		t.Errorf("error CreateConnectionPool: %s", err.Error())
		return
	}

	q.Component = "license"
	q.Action = "query"
	q.Item = "license"

	n := NewNsc(p)

	d, err = n.GetLicense()

	if err != nil {
		t.Errorf("error returned from GetLicense: %s", err.Error())
		return
	}

	if d.Status != "ok" || d.Result.Validity == "" {
		t.Errorf("expected status ok, and validity to be populated")
	}



}
