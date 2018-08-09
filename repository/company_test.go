package repository_test

import (
	"github.com/nosnibor89/irsales/model"

	"github.com/nosnibor89/irsales/repository"
	"github.com/nosnibor89/irsales/repository/mongo"

	"testing"
)

func TestInsert(t *testing.T) {
	MockDB := mongo.MockDB{}
	companyRepo := repository.CompanyDataStore{Database: MockDB}
	company := model.Company{
		Name: "Testing",
	}

	if err := companyRepo.Create(&company); err != nil {
		t.Error("Error testing CompanyRepository.Find", "Got"+company.Name)
	}
}
