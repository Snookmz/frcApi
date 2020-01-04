package database

import (
	"testing"
)

func TestCreateClass (t *testing.T) {
	db := PrepareTestDatabase()
	d := NewNsDatabase(db)

	var c Class
	c.NsId = 1
	c.Name = "unit test class"
	c.ParentId = 0
	c.PrevId = 15
	c.NextId = 17
	c.FirstChildId = 0
	c.LastChildId = 0
	c.Precedence = 9999
	c.PolicyId = 1

	n, err := d.CreateClass(c)

	if err != nil {
		t.Errorf("error returned from CreateClass: %s", err.Error())
		return
	}

	var classes []Class
	classes, err = d.GetClasses(c.PolicyId)

	if err != nil {
		t.Errorf("error returned from GetClasses: %s", err.Error())
		return
	}

	for _, j := range classes {
		if j.ClassId == n.ClassId {
			if j.NsId != c.NsId ||
				j.Name != c.Name ||
				j.PrevId != c.PrevId ||
				j.NextId != c.NextId ||
				j.FirstChildId != c.FirstChildId ||
				j.LastChildId != c.LastChildId ||
				j.Precedence != c.Precedence ||
				j.PolicyId != c.PolicyId {
				t.Errorf("expected new class to match what was sent")
			}
		}
	}

	err = d.DeleteClass(n.ClassId)

	if err != nil {
		t.Errorf("error returned from DeleteClass: %s\n", err.Error())
	}


}
