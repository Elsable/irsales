package repository

import (
	"log"

	"github.com/nosnibor89/irsales/repository/mongo"

	"github.com/nosnibor89/irsales/model"

	mgo "gopkg.in/mgo.v2"
)

const (
	// CompanyCollection is the name of the mongo db collection for companies
	CompanyCollection = "company"
)

type (

	// CompanyRepository acts as an interface for comopany datasource
	CompanyRepository interface {
		All() []model.Company
		Find(id int) (model.Company, error)
		Create(*model.Company) (model.Company, error)
		Update(*model.Company) (model.Company, error)
	}
)

// DB holds the mongodb session
var DB *mgo.Database

// Connect creates a session with a mongodb server and set a db in the repository. Return the session so anybody can close it later
func Connect() *mgo.Session {
	session, err := mgo.Dial("mongodb://admin:admin123@ds018508.mlab.com:18508/irsales")
	if err != nil {
		log.Println("Cound not dial to DB")
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	DB = session.DB("irsales")

	return mongo.MgoSession{Session: session}.Session
}
