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
	return []model.Company{}
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

func TestCreateCompany(t *testing.T) {
	cs := services.CompanyService{Repository: &mockRepository{}}

	newCompany := model.Company{Name: "deglotec", Location: "Santiago"}

	if _, err := cs.CreateCompany(&newCompany); err != nil {
		t.Error("Company.Create failed")
	}
}
