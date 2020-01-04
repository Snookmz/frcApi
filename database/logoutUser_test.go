package database

import (
	"gopkg.in/testfixtures.v2"
	"log"
	"testing"
)

func TestLogoutUser (t *testing.T) {
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

	_, err = d.LoginUser("heath", "n5admin")

	if err != nil {
		t.Errorf("error returned from LoginUser: %s", err.Error())
	}

	err = d.LogoutUser("heath")

	if err != nil {
		t.Errorf("error returned from LogoutUser: %s", err.Error())
	}

	err = d.LogoutUser("no user")
	if err == nil {
		t.Errorf("expected error")
	}



}
