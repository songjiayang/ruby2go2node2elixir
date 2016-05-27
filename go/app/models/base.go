package models

import (
	"time"

	"gopkg.in/mgo.v2"
)

var (
	db *mgo.Database
)

func init() {
	session, err := mgo.Dial("mongodb://localhost/ruby2go2node2elixir")
	if err != nil {
		panic(err)
	}

	if err := session.Ping(); err != nil {
		panic(err)
	}

	session.SetSafe(&mgo.Safe{
		W:        1,
		WTimeout: 200,
	})

	// set pool size
	session.SetPoolLimit(20)
	session.SetSyncTimeout(5 * time.Second)
	db = session.DB("ruby2go2node2elixir")
}
