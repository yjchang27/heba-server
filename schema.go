package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gorp.v1"
	"log"
	"time"
)

type User struct {
	Id          int64  `db:"user_id"`
	Nickname    string `db:"nickname"`
	Password    string `db:"password"`
	Name        string `db:"name"`
	Gender      bool   `db:"gender"`
	Age         int    `db:"age"`
	Affiliation string `db:"affiliation"`
	Created     int64  `db:"created"`
	LastLogin   int64  `db:"last_login"`
}

type Stamp struct {
	Id        int64   `db:"stamp_id"`
	PhotoLink string  `db:"photo_link"`
	Latitude  float64 `db:"latitude"`
	Longitude float64 `db:"longitude"`
	Captured  int64   `db:"captured"`
	Caption   string  `db:"caption"`
	EventId   int64   `db:"event_id"`
	Created   int64   `db:"created"`
}

type Event struct {
	Id          int64   `db:"event_id"`
	Name        string  `db:"name"`
	Explanation string  `db:"explanation"`
	Created     int64   `db:"created"`
	Latitude    float64 `db:"latitude"`
	Longitude   float64 `db:"longitude"`
}

func initDb() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	db, err := sql.Open("sqlite3", "temp.db")
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// add tables
	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "Id")
	dbmap.AddTableWithName(Stamp{}, "stamps").SetKeys(true, "Id")
	dbmap.AddTableWithName(Event{}, "events").SetKeys(true, "Id")

	// create the table.
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
}
