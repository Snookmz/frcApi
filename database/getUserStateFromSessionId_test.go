package database

import (
	"gopkg.in/testfixtures.v2"
	"log"
	"testing"
	"time"
)

func TestGetUserStateFromSessionId (t *testing.T) {
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

	var u UserState

	var sessionId string
	sessionId, err = d.LoginUser("heath", "n5admin")
	if err != nil {
		t.Errorf("error returned from LoginUser: %s", err.Error())
		return
	}

	u, err = d.GetUserStateFromSessionId(sessionId)

	if err != nil {
		t.Errorf("error returned from GetUserStateFromSessionId: %s", err.Error())
		return
	}

	if !u.SignedIn {
		t.Errorf("expected user state to be signed in")
	}

	var parsedTime time.Time
	parsedTime, err = time.Parse(time.RFC3339Nano, u.SignOnTime)
	if err != nil {
		t.Errorf("error returned from time.Parse: %s", err.Error())
		return
	}

	var timeNow time.Time
	timeNow = time.Now()

	year, month, day := timeNow.Date()
	year2, month2, day2 := parsedTime.Date()

	if year != year2 || month != month2 || day != day2 {
		t.Errorf("expected date to match")
	}

	if timeNow.Hour() != parsedTime.Hour() {
		t.Errorf("expected hour to match")
	}



}
