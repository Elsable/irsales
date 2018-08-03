package repository

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

// Connect creates a session with a mongodb server and set a db in the repository. Return the session so anybody can close it later
func Connect() *mgo.Session {
	session, err := mgo.Dial("mongodb://admin:admin123@ds018508.mlab.com:18508/irsales")
	if err != nil {
		log.Println("Cound not dial to DB")
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	db = session.DB("irsales")

	return session
}

// DB is a holder for the mongo db instance
func DB() *mgo.Database {
	return db
}
