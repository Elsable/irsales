package repository

import (
	"log"

	"github.com/nosnibor89/irsales/repository/mongo"

	"github.com/nosnibor89/irsales/model"

	mgo "gopkg.in/mgo.v2"
)

const (
	CompanyCollection = "Company"
)

type (

	// CompanyRepository acts as an interface for comopany datasource
	CompanyRepository interface {
		All() []model.Company
		Find(id int) (model.Company, error)
		Create(model.Company) error
		Update(model.Company) error
	}
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

	return mongo.MgoSession{Session: session}.Session
}

// DB is a holder for the mongo db instance
func DB() *mongo.MgoDatabase {
	return &mongo.MgoDatabase{Database: db}
}
