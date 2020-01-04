package database

import (
	"gopkg.in/testfixtures.v2"
	"log"
	"testing"
)

func TestGetUsers (t *testing.T) {

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

	var u User
	u.Name = "ns"
	u.GroupId = 2
	u.SiteId = 1
	u.Active = true

	var users []User
	users, err = d.GetUsers()
	if err != nil {
		t.Errorf("error returned from GetUsers: %s", err.Error())
		return
	}

	for _, u2 := range users {
		if u.Id == u2.Id {
			if u.Name != u2.Name ||
				u.GroupId != u2.GroupId ||
				u.SiteId != u2.SiteId ||
				u.Active != u2.Active {
				t.Errorf("expected User to have been edited")
			}
		}
	}

}
