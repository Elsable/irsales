package services_test

import (
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/nosnibor89/irsales/model"
	"github.com/nosnibor89/irsales/services"
)

type mockRepository struct{}

func (mr mockRepository) All() []model.Company {
	return []model.Company{
		model.Company{
			ID:       bson.NewObjectId(),
			Name:     "Tests company 1",
			Location: "Caracas",
		},
		model.Company{
			ID:       bson.NewObjectId(),
			Name:     "Tests company 2",
			Location: "Valencia",
		},
	}
}

func (mr mockRepository) Find(id int) (model.Company, error) {
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

	companies := cs.Find()
	count := len(companies)

	if count != 2 {
		t.Errorf("Companies not correcly fetched - got %v ,expected %v", count, 2)
	}
}
