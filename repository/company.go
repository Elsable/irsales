package repository

import (
	"github.com/nosnibor89/irsales/model"
	"github.com/nosnibor89/irsales/repository/mongo"
)

// CompanyDataStore holds the methods for getting company data from DB
type CompanyDataStore struct {
	Database mongo.DataLayer
}

// CompanyDataStoreWithDefault create a new company data store with the defeault DB in this package
func CompanyDataStoreWithDefault() *CompanyDataStore {
	return &CompanyDataStore{Database: DB()}
}

// NewCompanyDataStore create a new company data store with a provided DB
// func NewCompanyDataStore(dataBase *mongo.MongoDatabase) *CompanyDataStore {
// 	return &CompanyDataStore{dataBase}
// }

// All fetch companies from DB
func (dataStore *CompanyDataStore) All() ([]model.Company, error) {
	// testUser := struct {
	// 	Name, Lastname string
	// }{
	// 	Name:     "Robi",
	// 	Lastname: "Marquez",
	// }
	// db.C("users").Insert(&testUser)

	return []model.Company{}, nil
}

// Find company by id
func (dataStore *CompanyDataStore) Find(id int) (model.Company, error) {

	return model.Company{}, nil
}

// Create stores a company in DB
func (dataStore *CompanyDataStore) Create(company *model.Company) error {
	return nil
}

// Update store values for an existing company
func (dataStore *CompanyDataStore) Update(model.Company) error {
	return nil
}
