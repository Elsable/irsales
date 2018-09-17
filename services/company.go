package services

import (
	"github.com/nosnibor89/irsales/model"

	"github.com/nosnibor89/irsales/repository"
)

// CompanyServiceHandler is the interface for the company service
type CompanyServiceHandler interface {
	Create(company *model.Company) (model.Company, error)
	Find() []model.Company
	FindOne(id string) (model.Company, error)
}

// CompanyService holds the functionality for Companies features
type CompanyService struct {
	Repository repository.CompanyRepository
}

// NewCompanyService create a company service
func NewCompanyService() CompanyService {
	repo := repository.CompanyDataStoreWithDefault()
	return CompanyService{Repository: repo}
}

// NewCompanyServiceWith create a company service with a provided repo
func NewCompanyServiceWith(repo repository.CompanyRepository) CompanyService {
	return CompanyService{Repository: repo}
}

// Create creates a company
func (cs CompanyService) Create(company *model.Company) (model.Company, error) {

	created, err := cs.Repository.Create(company)
	if err != nil {
		return model.Company{}, err
	}

	return created, nil
}

// Find fetch companies
func (cs CompanyService) Find() []model.Company {
	companies := cs.Repository.All()
	return companies
}

// FindOne fetch a company using id provided
func (cs CompanyService) FindOne(id string) (model.Company, error) {
	found, err := cs.Repository.FindOne(id)
	if err != nil || string(found.ID) == "" {
		return model.Company{}, err
	}

	return found, nil
}
