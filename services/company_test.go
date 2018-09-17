package services_test

import (
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/nosnibor89/irsales/model"
	"github.com/nosnibor89/irsales/services"
	"github.com/nosnibor89/irsales/test"
)

type mockRepository struct{}

func (mr mockRepository) All() []model.Company {
	return test.MockCompanyData()
}

func (mr mockRepository) Find(id string) (model.Company, error) {
	return model.Company{}, nil
}

func (mr mockRepository) FindOne(id string) (model.Company, error) {
	for _, company := range test.MockCompanyData() {
		if company.ID == bson.ObjectId(id) {
			return company, nil
		}
	}
	return model.Company{}, nil
}

func (mr mockRepository) Create(c *model.Company) (model.Company, error) {
	c.ID = bson.NewObjectId()
	return *c, nil
}

func (mr mockRepository) Update(c *model.Company) (model.Company, error) {
	c.UpdatedAt = time.Now().String()
	return *c, nil
}

func TestCreate(t *testing.T) {
	cs := services.CompanyService{Repository: &mockRepository{}}

	newCompany := model.Company{Name: "deglotec", Location: "Santiago"}

	if _, err := cs.Create(&newCompany); err != nil {
		t.Error("Company.Create failed")
	}
}

func TestAll(t *testing.T) {
	cs := services.CompanyService{Repository: &mockRepository{}}

	companyData := test.MockCompanyData()
	expectedCompany := companyData[0]
	id := string(expectedCompany.ID)

	if company, err := cs.FindOne(id); company.ID == "" || err != nil {
		t.Errorf("Company not correctly fetched - got %v ,expected %v", company.ID, expectedCompany.ID)
	}
}

func TestFind(t *testing.T) {
	cs := services.CompanyService{Repository: &mockRepository{}}

	companies := cs.Find()
	count := len(companies)

	if count != 2 {
		t.Errorf("Companies not correctly fetched - got %v ,expected %v", count, 2)
	}
}
