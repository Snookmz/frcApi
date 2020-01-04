package database

import "testing"

func TestDeleteClass (t *testing.T) {
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
	}


	err = d.DeleteClass(n.ClassId)

	if err != nil {
		t.Errorf("error returned from DeleteClass: %s\n", err.Error())
	}

	var classes []Class
	classes, err = d.GetClasses(c.PolicyId)

	for _, j := range classes {
		if j.ClassId == n.ClassId {
			t.Errorf("expected class to be deleted")
		}
	}



}
