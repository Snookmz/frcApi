package database

import (
	"testing"
)

func TestGetClasses (t *testing.T) {
	var err error
	var classes []Class
	var d *NsDatabase

	db := PrepareTestDatabase()
	d = NewNsDatabase(db)

	classes, err = d.GetClasses(1)

	if err != nil {
		t.Errorf("error returned from GetPolicyClasses: %s", err.Error())
	}

	if len(classes) == 0 {
		t.Errorf("expected classes to be populated")
	}

	for _, c := range classes {
		if c.ClassId == 0 {
			t.Errorf("expected classId to be non-zero")
		}
		if c.Name == "" {
			t.Errorf("expected Name to be populated")
		}
	}



}