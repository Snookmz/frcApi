package database

import (
	"gopkg.in/testfixtures.v2"
	"log"
	"testing"
)

func TestEditClass (t *testing.T) {
	db := PrepareTestDatabase()
	d := NewNsDatabase(db)

	var fixtures *testfixtures.Context
	var err error
	fixtures, err = testfixtures.NewFolder(db, &testfixtures.PostgreSQL{}, "./fixtures")
	if err != nil {
		log.Printf("error returned from textfixtures.NewFolder: %s", err.Error())
		log.Fatal(err)
	}

	err = fixtures.Load()
	if err != nil {
		log.Printf("error returned from fixtures.Load: %s\n", err.Error())
		log.Fatal(err)
	}

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

	var n Class
	n, err = d.CreateClass(c)

	if err != nil {
		t.Errorf("error returned from CreateClass: %s", err.Error())
		return
	}

	n.Name = "unit test class edited"
	n.ParentId = 1
	n.PrevId = 16
	n.NextId = 18
	n.FirstChildId = 1
	n.LastChildId = 1
	n.Precedence = 9998
	n.PolicyId = 2

	err = d.EditClass(n)
	if err != nil {
		t.Errorf("error returned from EditClass: %s", err.Error())
		return
	}

	var classes []Class
	classes, err = d.GetClasses(n.PolicyId)

	for _, j := range classes {
		if j.ClassId == n.ClassId {
			if j.NsId != n.NsId ||
				j.Name != n.Name ||
				j.PrevId != n.PrevId ||
				j.NextId != n.NextId ||
				j.FirstChildId != n.FirstChildId ||
				j.LastChildId != n.LastChildId ||
				j.Precedence != n.Precedence ||
				j.PolicyId != n.PolicyId {
				t.Errorf("expected class to match what was sent")
			}
		}
	}

	err = d.DeleteClass(n.ClassId)

	if err != nil {
		t.Errorf("error returned from DeleteClass: %s\n", err.Error())
	}

}