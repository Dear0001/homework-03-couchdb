package db

import (
	kivik "github.com/go-kivik/kivik/v4"
	_ "github.com/go-kivik/kivik/v4/couchdb"
)

var db *kivik.DB

// InitCouchDB initializes the CouchDB client and database instance
func InitCouchDB() *kivik.DB {
	client, err := kivik.New("couch", "http://admin:123@localhost:5985/")
	if err != nil {
		panic(err)
	}
	db = client.DB("students")
	return db
}

func GetDB() *kivik.DB {
	if db == nil {
		InitCouchDB()
	}
	return db
}
