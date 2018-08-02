package config

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

func SetDB() {
	session, err := mgo.Dial("mongodb://admin:admin123@ds018508.mlab.com:18508/irsales")
	if err != nil {
		log.Println("Cound not dial to DB")
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	db = session.DB("irsales")

	testUser := struct {
		Name, Lastname string
	}{
		Name:     "Robi",
		Lastname: "Marquez",
	}
	db.C("users").Insert(&testUser)
}

func DB() *mgo.Database {
	return db
}
