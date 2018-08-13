package repository

import (
	"time"

	"github.com/nosnibor89/irsales/model"
	"github.com/nosnibor89/irsales/repository/mongo"
	"gopkg.in/mgo.v2/bson"
)

// CompanyDataStore holds the methods for getting company data from DB
type CompanyDataStore struct {
	Database mongo.DataLayer
}

// CompanyDataStoreWithDefault create a new company data store with the defeault DB in this package
func CompanyDataStoreWithDefault() *CompanyDataStore {
	return &CompanyDataStore{Database: mongo.MgoDatabase{Database: DB}}
}

// All fetch companies from DB
func (dataStore *CompanyDataStore) All() []model.Company {
	return []model.Company{}
}

// Find company by id
func (dataStore *CompanyDataStore) Find(id int) (model.Company, error) {
	return model.Company{}, nil
}

// Create stores a company in DB
func (dataStore *CompanyDataStore) Create(company *model.Company) (model.Company, error) {
	c := dataStore.Database.C(CompanyCollection)

	company.ID = bson.NewObjectId()
	company.CreatedAt = time.Now().String()
	if err := c.Insert(company); err != nil {
		return model.Company{}, err
	}

	return *company, nil
}

// Update store values for an existing company
func (dataStore CompanyDataStore) Update(company *model.Company) (model.Company, error) {
	return model.Company{}, nil
}
